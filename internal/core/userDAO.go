package core

import "github.com/google/uuid"

type UserDAO struct {
	Id            uuid.UUID `db:"id"`
	Gmail         string    `db:"gmail"`
	Username      string    `db:"username"`
	Nickname      string    `db:"nickname"`
	IsRegistered  bool      `db:"is_registered"`
	HasProfilePic bool      `db:"has_picture"`
}

func (u *UserDAO) ToDomain() (user User) {
	user.Id = u.Id
	user.Gmail = u.Gmail
	//if u.Username == nil {
	//	user.Username = ""
	//} else {
	//	user.Username = *u.Username
	//}
	user.Username = u.Username
	user.Nickname = u.Nickname
	user.IsRegistered = u.IsRegistered
	user.HasProfilePic = u.HasProfilePic
	return
}
