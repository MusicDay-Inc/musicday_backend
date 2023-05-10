package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"server/internal/core"
)

type ReviewRepository struct {
	db *sqlx.DB
}

func (r ReviewRepository) GetReviewsFromUser(userId uuid.UUID, limit int, offset int) (reviews []core.ReviewDAO, err error) {
	q := `
	SELECT * FROM reviews WHERE user_id = $1
	ORDER BY published_at DESC 
	LIMIT $2 OFFSET $3
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Select(&reviews, q, userId, limit, offset)
	if err != nil {
		return reviews, err
	}
	return
}

func (r ReviewRepository) GetAlbumReviewsFromUser(userId uuid.UUID, limit int, offset int) (reviews []core.ReviewDAO, err error) {
	q := `
	SELECT * FROM reviews WHERE (user_id, is_song_reviewed) = ($1, false)
	ORDER BY published_at DESC
	LIMIT $2 OFFSET $3
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Select(&reviews, q, userId, limit, offset)
	if err != nil {
		return reviews, err
	}
	return
}

func (r ReviewRepository) GetSongReviewsFromUser(userId uuid.UUID, limit int, offset int) (reviews []core.ReviewDAO, err error) {
	q := `
	SELECT * FROM reviews WHERE (user_id, is_song_reviewed) = ($1, true)
	ORDER BY published_at DESC
	LIMIT $2 OFFSET $3
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Select(&reviews, q, userId, limit, offset)
	if err != nil {
		return reviews, err
	}
	return
}

func (r ReviewRepository) Delete(id uuid.UUID) error {
	q := `
	DELETE FROM reviews  WHERE id = $1;
	`
	logrus.Trace(formatQuery(q))
	_, err := r.db.Exec(q, id)
	return err
}

func (r ReviewRepository) GetSubscriptionReviews(releaseId uuid.UUID, userId uuid.UUID, limit int, offset int) (reviews []core.ReviewDAO, err error) {
	q := `
	SELECT id, user_id, is_song_reviewed, release_id, published_at, score, review_text
	FROM reviews
         JOIN subscriptions on subscriptions.subscriber_id = $1
	WHERE (subscriptions.subscription_id = reviews.user_id 
	           AND reviews.release_id = $2)
	ORDER BY published_at
	LIMIT $3 OFFSET $4;
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Select(&reviews, q, userId, releaseId, limit, offset)
	if err != nil {
		logrus.Error(err)
		return reviews, err
	}
	return reviews, nil
}

func (r ReviewRepository) UpdateReview(review core.Review) (res core.ReviewDAO, err error) {
	q := `
	UPDATE reviews
	SET (published_at,
     	score,
     	review_text) 
	    = ($1, $2, $3)
	WHERE (user_id, release_id) = ($4, $5)
	RETURNING *
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&res, q, review.PublishedAt, review.Score, review.Text, review.UserId, review.ReleaseId)
	if err != nil {
		logrus.Error(err)
	}
	return
}
func (r ReviewRepository) ExistsFromUser(userId uuid.UUID, reviewId uuid.UUID) (res bool, err error) {
	q := `
	SELECT EXISTS(SELECT
	FROM reviews
	WHERE (id, user_id)= ($1, $2));
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&res, q, reviewId, userId)
	if err != nil {
		return false, err
	}
	return res, err
}

func (r ReviewRepository) ExistsToRelease(userId uuid.UUID, releaseId uuid.UUID) (res bool, err error) {
	q := `
	SELECT EXISTS(SELECT
	FROM reviews
	WHERE (user_id, release_id)= ($1, $2));
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&res, q, userId, releaseId)
	if err != nil {
		return false, err
	}
	return res, err
}

func (r ReviewRepository) InsertReview(review core.Review) (res core.ReviewDAO, err error) {
	q := `
	INSERT INTO reviews 
	    (user_id, 
	     is_song_reviewed,
	     release_id,
	     published_at, 
	     score,
	     review_text)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING *
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&res, q,
		review.UserId,
		review.IsSongReviewed,
		review.ReleaseId,
		review.PublishedAt,
		review.Score,
		review.Text)
	if err != nil {
		logrus.Error(err)
	}
	return
}

func (r ReviewRepository) GetById(id uuid.UUID) (review core.ReviewDAO, err error) {
	q := `
	SELECT * FROM reviews WHERE id = $1
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&review, q, id)
	if err != nil {
		return review, err
	}
	return
}

func (r ReviewRepository) GetReviewToRelease(releaseId uuid.UUID, userId uuid.UUID) (review core.ReviewDAO, err error) {
	q := `
	SELECT * FROM reviews WHERE (user_id, release_id) = ($1, $2)
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&review, q, userId, releaseId)
	if err != nil {
		return review, err
	}
	return
}

func NewReviewRepository(db *sqlx.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}
