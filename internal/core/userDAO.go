package core

import (
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
	SubscriberAmount   int       `db:"subscribers_c"`
	SubscriptionAmount int       `db:"subscriptions_c"`
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

type ReviewDAO struct {
	Id             uuid.UUID `db:"id"`
	UserId         uuid.UUID `db:"user_id"`
	IsSongReviewed bool      `db:"is_song_reviewed"`
	ReleaseId      uuid.UUID `db:"release_id"`
	PublishedAt    time.Time `db:"published_at"`
	Score          uint8     `db:"score"`
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
