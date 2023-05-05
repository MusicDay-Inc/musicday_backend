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

func (r ReviewRepository) Exists(userId uuid.UUID, releaseId uuid.UUID) (res bool, err error) {
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
	//row := r.db.QueryRow(q,
	//	review.Id,
	//	review.UserId,
	//	review.IsSongReviewed,
	//	review.ReleaseId,
	//	review.PublishedAt,
	//	review.Score,
	//	review.Text,
	//)
	//err = row.Scan(&res)
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
