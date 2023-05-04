package core

import "github.com/google/uuid"

type UserDTO struct {
	Id                 uuid.UUID `json:"id,omitempty"`
	Gmail              string    `json:"gmail,omitempty"`
	Username           string    `json:"username"`
	Nickname           string    `json:"nickname"`
	IsRegistered       bool      `json:"is_registered"`
	HasProfilePic      bool      `json:"has_picture"`
	SubscriberAmount   int       `json:"subscriber_amount"`
	SubscriptionAmount int       `json:"subscription_amount"`
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

type ReviewDTO struct {
	Id             uuid.UUID `json:"id,omitempty"`
	UserId         uuid.UUID `json:"user_id,omitempty"`
	IsSongReviewed bool      `json:"is_song_reviewed,omitempty"`
	Song           Song      `json:"song,omitempty"`
	Album          Album     `json:"album,omitempty"`
	PublishedAt    bool      `json:"published_at,omitempty"`
	Score          uint8     `json:"score,omitempty"`
	Text           string    `json:"review_text,omitempty"`
}
