package core

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
	Id            int    `json:"user_id,omitempty"`
	Gmail         string `json:"gmail,omitempty"`
	Nickname      string `json:"nickname" binding:"required"`
	Username      string `json:"username" binding:"required"`
	IsRegistered  bool   `json:"is_registered"`
	HasProfilePic bool   `json:"has_profile_pic,omitempty"`
}

//func (u *User) ToDTO(user UserDTO) {
//	user.Id = strconv.Itoa(u.Id)
//	user.Gmail = u.Gmail
//	user.Nickname = u.Nickname
//	user.Username = u.Username
//	user.IsRegistered = u.IsRegistered
//	user.HasProfilePic = u.HasProfilePic
//}
