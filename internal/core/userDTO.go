package core

type UserDTO struct {
	Id            string `json:"user_id,omitempty"`
	Gmail         string `json:"gmail,omitempty"`
	Nickname      string `json:"nickname" binding:"required"`
	Username      string `json:"username" binding:"required"`
	IsRegistered  bool   `json:"is_registered"`
	HasProfilePic bool   `json:"has_profile_pic,omitempty"`
}
