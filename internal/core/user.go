package core

type User struct {
	Id       int    `json:"id"`
	Gmail    string `json:"gmail"`
	Nickname string `json:"nickname" binding:"required"`
	Username string `json:"username" binding:"required"`
	//AvatarId string `json:"avatar_id" binding:"required"`
	IsRegistered  bool `json:"is_registered"`
	HasProfilePic bool `json:"has_profile_pic"`
}
