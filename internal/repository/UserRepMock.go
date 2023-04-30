package repository

import (
	"errors"
	"server/internal/core"
	"strconv"
)

type UserRepo struct {
	users []core.User
}

func (rep *UserRepo) GetByUsername(username string) (core.User, error) {
	ind := findUsername(rep.users, username)
	if ind == len(rep.users) {
		return core.User{}, errors.New("username doesn't exist")
	}
	return rep.users[ind], nil
}

func NewUserRepo() *UserRepo {
	data := []core.User{{0, "mock@gmail.ru", "Mock", "mock", true, false}}
	return &UserRepo{users: data}
}

func (rep *UserRepo) Create(gmail string) (int, error) {
	user := core.User{Id: len(rep.users), Gmail: gmail, Nickname: "NewUser", Username: "mock" + strconv.Itoa(len(rep.users)), IsRegistered: false, HasProfilePic: false}
	rep.users = append(rep.users, user)
	return user.Id, nil
}

func (rep *UserRepo) Exists(gmail string) bool {
	ind := findGmail(rep.users, gmail)
	return ind < len(rep.users)
}

func (rep *UserRepo) GetById(userId int) (core.User, error) {
	//err := ((releaseId < 0 || releaseId > len(rep.users)  "error": "ok")
	//var err error
	var item core.User
	if userId < 0 || userId >= len(rep.users) {
		err := errors.New("incorrect  UserID")
		return item, err
	}
	err := error(nil)
	return rep.users[userId], err
}

func (rep *UserRepo) GetByGmail(gmail string) (core.User, error) {
	//err := ((releaseId < 0 || releaseId > len(rep.users)  "error": "ok")
	//var err error
	var item core.User
	ind := findGmail(rep.users, gmail)
	if ind >= len(rep.users) {
		err := errors.New("incorrect  UserID")
		return item, err
	}
	err := error(nil)
	return rep.users[ind], err
}

func (rep *UserRepo) Change(userId int, u core.User) core.User {
	rep.users[userId].Username = u.Username
	rep.users[userId].Nickname = u.Nickname
	rep.users[userId].IsRegistered = true
	rep.users[userId].HasProfilePic = u.HasProfilePic
	return rep.users[userId]
}

func findGmail(users []core.User, gmail string) int {
	for i, u := range users {
		if gmail == u.Gmail {
			return i
		}
	}
	return len(users)
}
func findUsername(users []core.User, username string) int {
	for i, u := range users {
		if username == u.Username {
			return i
		}
	}
	return len(users)
}
