package core

import (
	"github.com/google/uuid"
	"time"
)

type UserDTO struct {
	Id                 uuid.UUID `json:"id,omitempty"`
	Gmail              string    `json:"gmail,omitempty"`
	Username           string    `json:"username"`
	Nickname           string    `json:"nickname"`
	IsRegistered       bool      `json:"is_registered"`
	HasProfilePic      bool      `json:"has_picture"`
	SubscriberAmount   int32     `json:"subscriber_amount"`
	SubscriptionAmount int32     `json:"subscription_amount"`
}

func (u *UserDTO) ToDomain() (user User) {
	user.Id = u.Id
	user.Gmail = u.Gmail
	user.Nickname = u.Nickname
	user.Username = u.Username
	user.IsRegistered = u.IsRegistered
	user.HasProfilePic = u.HasProfilePic
	user.SubscriberAmount = u.SubscriberAmount
	user.SubscriptionAmount = u.SubscriptionAmount
	return
}

type SearchDTO struct {
	Request string `json:"search,omitempty" binding:"required"`
}

type ReviewDTO struct {
	Id             uuid.UUID `json:"id,omitempty"`
	UserId         uuid.UUID `json:"user_id,omitempty"`
	IsSongReviewed bool      `json:"is_song_reviewed,omitempty"`
	ReleaseId      uuid.UUID `json:"release_id,omitempty" binding:"required"`
	Song           Song      `json:"song,omitempty"`
	Album          Album     `json:"album,omitempty"`
	PublishedAt    time.Time `json:"published_at,omitempty"`
	Score          int32     `json:"score,omitempty" binding:"required"`
	Text           string    `json:"review_text,omitempty"`
}

// FormReview forms review (first param) from user (second param)
func (r ReviewDTO) FormReview(releaseId uuid.UUID, userId uuid.UUID) (review Review) {
	review.Id = r.Id
	review.ReleaseId = releaseId
	review.UserId = userId
	review.PublishedAt = time.Now()
	review.Score = r.Score
	review.Text = r.Text
	return
}
