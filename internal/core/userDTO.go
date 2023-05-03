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
