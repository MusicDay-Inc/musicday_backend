package core

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname" binding:"required"`
	Username string `json:"username" binding:"required"`
	//AvatarId string `json:"avatar_id" binding:"required"`
	HasProfilePic bool `json:"has_profile_pic" binding:"required"`
	//GoogleId int    `json:"google_id"`
}
