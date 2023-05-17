package repository

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"server/internal/core"
)

type UserRepository struct {
	db *sqlx.DB
}

func (r UserRepository) GetPlayerID(userId uuid.UUID) (res string, err error) {
	q := `
	SELECT app_id FROM user_appid WHERE user_id = $1
	`
	err = r.db.Get(&res, q, userId)
	if err != nil {
		logrus.Error(err)
		return
	}
	return
}

func (r UserRepository) InstallAppID(clientId uuid.UUID, playerID uuid.UUID) error {
	q := `
	INSERT INTO user_appid VALUES 
	 ($1, $2)
	`
	_, err := r.db.Exec(q, clientId, playerID)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r UserRepository) GetBio(userId uuid.UUID) (res string, err error) {
	q := `
	SELECT bio
	FROM user_bios
	WHERE user_id = $1;
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&res, q, userId)
	return res, err
}

func (r UserRepository) CreateBio(userId uuid.UUID, bio string) (string, error) {
	q := `
	INSERT INTO user_bios (user_id, bio) VALUES ($1, $2) 
	`
	logrus.Trace(formatQuery(q))
	_, err := r.db.Exec(q, userId, bio)
	if err != nil {
		return "", err
	}
	return bio, nil
}

func (r UserRepository) GetSubscriptionsOf(userId uuid.UUID, limit int, offset int) (users []core.UserDAO, err error) {
	q := `
	SELECT id, gmail, username, nickname, is_registered, has_picture, subscribers_c, subscriptions_c
	FROM users
         INNER JOIN subscriptions on (users.id = subscriptions.subscription_id AND
                                      subscriptions.subscriber_id = $1)
	ORDER BY username 
	LIMIT $2 OFFSET $3;
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Select(&users, q, userId, limit, offset)
	return
}

func (r UserRepository) GetSubscribers(userId uuid.UUID, limit int, offset int) (users []core.UserDAO, err error) {
	q := `
	SELECT id, gmail, username, nickname, is_registered, has_picture, subscribers_c, subscriptions_c
	FROM users
         INNER JOIN subscriptions on (users.id = subscriptions.subscriber_id AND
                                      subscriptions.subscription_id = $1)
	ORDER BY username 
	LIMIT $2 OFFSET $3;
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Select(&users, q, userId, limit, offset)
	return
}

func (r UserRepository) IsSubscriptionExists(clientId uuid.UUID, userId uuid.UUID) (res bool) {
	q := `
	SELECT EXISTS(SELECT
	FROM subscriptions
	WHERE (subscriber_id, subscription_id) = ($1, $2));
	`
	logrus.Trace(formatQuery(q))
	err := r.db.Get(&res, q, clientId, userId)
	if err != nil {
		return false
	}
	return res
}

func (r UserRepository) ExistsWithId(id uuid.UUID) (res bool) {
	q := `
	SELECT EXISTS(SELECT
	FROM users
	WHERE id = $1);
	`
	logrus.Trace(formatQuery(q))
	err := r.db.Get(&res, q, id)
	if err != nil {
		return false
	}
	return res
}

func (r UserRepository) SearchUsers(query string, clientId uuid.UUID, limit int, offset int) (users []core.UserDAO, err error) {
	q := `
	SELECT *
		FROM users
	WHERE users.username ILIKE $1 || '%'
	ORDER BY username
	LIMIT $2 OFFSET $3;
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Select(&users, q, query, limit, offset)
	return
}

func (r UserRepository) Subscribe(clientId uuid.UUID, userId uuid.UUID) (core.User, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return core.User{}, err
	}
	var (
		user core.UserDAO
	)
	q := `
	INSERT INTO subscriptions VALUES 
	 ($1, $2)
	`
	logrus.Trace(formatQuery(q))
	_, err = tx.Exec(q, clientId, userId)
	if err != nil {
		errRollback := tx.Rollback()
		logrus.Error(errRollback)
		return core.User{}, err
	}
	q = `
	UPDATE users 
   	SET subscriptions_c = subscriptions_c + 1
		WHERE id = $1;
	`
	_, err = tx.Exec(q, clientId)
	if err != nil {
		errRollback := tx.Rollback()
		logrus.Error(errRollback)
		return core.User{}, err
	}
	q = `
	UPDATE users 
   	SET subscribers_c = subscribers_c + 1
		WHERE id = $1
	RETURNING id, gmail, username, nickname, is_registered, has_picture, subscribers_c, subscriptions_c;
	`
	logrus.Trace(formatQuery(q))
	row := tx.QueryRow(q, userId)
	err = row.Scan(&user.Id, &user.Gmail, &user.Username, &user.Nickname,
		&user.IsRegistered, &user.HasProfilePic, &user.SubscriberAmount, &user.SubscriptionAmount)
	if err != nil {
		logrus.Error(err)
		tx.Rollback()
		return core.User{}, err
	}
	return user.ToDomain(), tx.Commit()
}

func (r UserRepository) Unsubscribe(clientId uuid.UUID, userId uuid.UUID) (core.User, error) {
	var (
		user core.UserDAO
	)
	tx, err := r.db.Begin()
	if err != nil {
		return core.User{}, err
	}
	q := `
	DELETE FROM subscriptions WHERE (subscriber_id, subscription_id) = ($1, $2) 
	`
	logrus.Trace(formatQuery(q))
	_, err = tx.Exec(q, clientId, userId)
	if err != nil {
		errRollback := tx.Rollback()
		logrus.Error(errRollback)
		return core.User{}, err
	}
	q = `
	UPDATE users 
   	SET subscriptions_c = subscriptions_c - 1
		WHERE id = $1;
	`
	_, err = tx.Exec(q, clientId)
	if err != nil {
		errRollback := tx.Rollback()
		logrus.Error(errRollback)
		return core.User{}, err
	}
	q = `
	UPDATE users 
   	SET subscribers_c = subscribers_c - 1
		WHERE id = $1
	RETURNING id, gmail, username, nickname, is_registered, has_picture, subscribers_c, subscriptions_c;
	`
	logrus.Trace(formatQuery(q))
	row := tx.QueryRow(q, userId)
	err = row.Scan(&user.Id, &user.Gmail, &user.Username, &user.Nickname,
		&user.IsRegistered, &user.HasProfilePic, &user.SubscriberAmount, &user.SubscriptionAmount)
	if err != nil {
		logrus.Error(err)
		tx.Rollback()
		return core.User{}, err
	}
	return user.ToDomain(), tx.Commit()
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r UserRepository) Create(gmail string) (uuid.UUID, error) {
	q := `
	INSERT INTO users (gmail, is_registered)
	VALUES ($1, false)
	RETURNING id
	`
	logrus.Trace(formatQuery(q))
	row := r.db.QueryRow(q, gmail)
	var (
		id  uuid.UUID
		err error
	)
	err = row.Scan(&id)
	return id, err
}

func (r UserRepository) Exists(gmail string) (res bool) {
	q := `
	SELECT EXISTS(SELECT
	FROM users
	WHERE gmail = $1);
	`
	logrus.Trace(formatQuery(q))
	err := r.db.Get(&res, q, gmail)
	if err != nil {
		return false
	}
	return res
}

func (r UserRepository) GetById(userId uuid.UUID) (core.User, error) {
	q := `
	SELECT * FROM users WHERE id = $1
	`
	logrus.Trace(formatQuery(q))
	var user core.UserNullableDAO
	err := r.db.Get(&user, q, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return core.User{}, core.ErrNotFound
		}
		return core.User{}, err
	}
	return user.ToDomain(), nil
}

func (r UserRepository) GetByUsername(username string) (user core.UserDAO, err error) {
	q := `
	SELECT * FROM users WHERE username = $1
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&user, q, username)
	if err != nil {
		return core.UserDAO{}, err
	}
	return user, nil
}

func (r UserRepository) GetByGmail(gmail string) (core.User, error) {
	q := `
	SELECT * FROM users WHERE gmail = $1
	`
	logrus.Trace(formatQuery(q))
	var user core.UserNullableDAO
	err := r.db.Get(&user, q, gmail)
	//row := r.db.QueryRow(q, gmail)
	//err = row.Scan(&user)
	if err != nil {
		return core.User{}, err
	}
	return user.ToDomain(), nil
}

func (r UserRepository) Register(u core.User) (user core.UserDAO, err error) {
	q := `
	UPDATE users
	SET (username, nickname, is_registered) = ($1, $2, true)
	WHERE id = $3
	RETURNING *
	`
	logrus.Trace(formatQuery(q))
	//row := r.db.QueryRow(q, u.Username, u.Nickname, u.Id)
	//err = row.Scan(&user)
	err = r.db.Get(&user, q, u.Username, u.Nickname, u.Id)
	return user, err
}

func (r UserRepository) ChangeUsername(userId uuid.UUID, username string) (user core.UserDAO, err error) {
	q := `
	UPDATE users
	SET username = $1
	WHERE id = $2
	RETURNING *
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&user, q, username, userId)
	return
}
func (r UserRepository) ChangeNickname(userId uuid.UUID, nickname string) (user core.UserDAO, err error) {
	q := `
	UPDATE users
	SET nickname = $1
	WHERE id = $2
	RETURNING *
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&user, q, nickname, userId)
	return
}
func (r UserRepository) InstallPicture(id uuid.UUID) (user core.UserDAO, err error) {
	q := `
	UPDATE users
	SET has_picture = true
	WHERE id = $1
	RETURNING *
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&user, q, id)
	return
}
