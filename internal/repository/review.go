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
	SELECT * FROM reviews WHERE (song_or_album_id, user_id) = ($1, $2)
	`
	logrus.Trace(formatQuery(q))
	err = r.db.Get(&review, q, releaseId, userId)
	if err != nil {
		return review, err
	}
	return
}

func NewReviewRepository(db *sqlx.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}
