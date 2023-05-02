package core

// TODO
type UserDAO struct {
	Id            int    `db:"user_id"`
	Gmail         string `db:"gmail"`
	Nickname      string `db:"nickname"`
	Username      string `db:"username"`
	IsRegistered  bool   `db:"is_registered"`
	HasProfilePic bool   `db:"has_profile_pic"`
}

func (user *UserDAO) ToDomain() {
	// TODO
}
