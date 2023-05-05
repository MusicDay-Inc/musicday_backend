package core

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type UserDAO struct {
	Id                 uuid.UUID `db:"id"`
	Gmail              string    `db:"gmail"`
	Username           string    `db:"username"`
	Nickname           string    `db:"nickname"`
	IsRegistered       bool      `db:"is_registered"`
	HasProfilePic      bool      `db:"has_picture"`
	SubscriberAmount   int32     `db:"subscribers_c"`
	SubscriptionAmount int32     `db:"subscriptions_c"`
}
type UserNullableDAO struct {
	Id                 uuid.UUID      `db:"id"`
	Gmail              string         `db:"gmail"`
	Username           sql.NullString `db:"username"`
	Nickname           sql.NullString `db:"nickname"`
	IsRegistered       bool           `db:"is_registered"`
	HasProfilePic      bool           `db:"has_picture"`
	SubscriberAmount   sql.NullInt32  `db:"subscribers_c"`
	SubscriptionAmount sql.NullInt32  `db:"subscriptions_c"`
}

func (u *UserDAO) ToDomain() (user User) {
	user.Id = u.Id
	user.Gmail = u.Gmail
	user.Username = u.Username
	user.Nickname = u.Nickname
	user.IsRegistered = u.IsRegistered
	user.HasProfilePic = u.HasProfilePic
	user.SubscriberAmount = u.SubscriberAmount
	user.SubscriptionAmount = u.SubscriptionAmount
	return
}
func (u *UserNullableDAO) ToDomain() (user User) {
	user.Id = u.Id
	user.Gmail = u.Gmail
	if u.Username.Valid {
		user.Username = u.Username.String
	}
	if u.Nickname.Valid {
		user.Nickname = u.Nickname.String
	}
	user.IsRegistered = u.IsRegistered
	user.HasProfilePic = u.HasProfilePic
	if u.SubscriberAmount.Valid {
		user.SubscriberAmount = u.SubscriberAmount.Int32
	}
	if u.SubscriptionAmount.Valid {
		user.SubscriptionAmount = u.SubscriptionAmount.Int32
	}
	return
}

type ReviewDAO struct {
	Id             uuid.UUID `db:"id"`
	UserId         uuid.UUID `db:"user_id"`
	IsSongReviewed bool      `db:"is_song_reviewed"`
	ReleaseId      uuid.UUID `db:"release_id"`
	PublishedAt    time.Time `db:"published_at"`
	Score          int32     `db:"score"`
	Text           string    `db:"review_text"`
}

func (r *ReviewDAO) ToDomain() (review Review) {
	review.Id = r.Id
	review.UserId = r.UserId
	review.IsSongReviewed = r.IsSongReviewed
	review.ReleaseId = r.ReleaseId
	review.PublishedAt = r.PublishedAt
	review.Score = r.Score
	review.Text = r.Text
	return
}

type ReviewNullableDAO struct {
	Id             uuid.UUID      `db:"id"`
	UserId         uuid.UUID      `db:"user_id"`
	IsSongReviewed sql.NullBool   `db:"is_song_reviewed"`
	ReleaseId      uuid.UUID      `db:"release_id"`
	PublishedAt    sql.NullTime   `db:"published_at"`
	Score          sql.NullInt32  `db:"score"`
	Text           sql.NullString `db:"review_text"`
}

func (r *ReviewNullableDAO) ToDomain() Review {
	var review Review
	review.Id = r.Id
	review.UserId = r.UserId
	if r.IsSongReviewed.Valid {
		review.IsSongReviewed = r.IsSongReviewed.Bool
	}
	review.ReleaseId = r.ReleaseId
	if r.PublishedAt.Valid {
		review.PublishedAt = r.PublishedAt.Time
	}
	if r.Score.Valid {
		review.Score = r.Score.Int32
	}
	if r.Text.Valid {
		review.Text = r.Text.String
	}
	return review
}
