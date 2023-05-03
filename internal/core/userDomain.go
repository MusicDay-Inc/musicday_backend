package core

import (
	"github.com/google/uuid"
)

// GOOD?
//type User struct {
//	Id            int
//	Gmail         string
//	Nickname      string
//	Username      string
//	IsRegistered  bool
//	HasProfilePic bool
//}

type User struct {
	Id            uuid.UUID `json:"user_id,omitempty"`
	Gmail         string    `json:"gmail,omitempty"`
	Username      string    `json:"username" binding:"required"`
	Nickname      string    `json:"nickname" binding:"required"`
	IsRegistered  bool      `json:"is_registered"`
	HasProfilePic bool      `json:"has_profile_pic,omitempty"`
}

func (u *User) ToDTO() (user UserDTO) {
	//user.Id = u.Id.String()
	user.Id = u.Id
	user.Gmail = u.Gmail
	user.Nickname = u.Nickname
	user.Username = u.Username
	user.IsRegistered = u.IsRegistered
	user.HasProfilePic = u.HasProfilePic
	return
}
