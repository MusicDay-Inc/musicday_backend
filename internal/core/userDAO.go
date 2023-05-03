package core

import "github.com/google/uuid"

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
