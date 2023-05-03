package core

import (
	"github.com/google/uuid"
)

type User struct {
	Id                 uuid.UUID
	Gmail              string
	Username           string
	Nickname           string
	IsRegistered       bool
	HasProfilePic      bool
	SubscriberAmount   int
	SubscriptionAmount int
}

func (u *User) ToDTO() (user UserDTO) {
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
