package repository

import (
	"errors"
	"github.com/google/uuid"
	"server/internal/core"
)

type UserMockRepository struct {
	users []core.User
}

func (rep *UserMockRepository) GetByUsername(username string) (core.User, error) {
	ind := findUsername(rep.users, username)
	if ind == len(rep.users) {
		return core.User{}, errors.New("username doesn't exist")
	}
	return rep.users[ind], nil
}

func NewUserMockRepo() *UserMockRepository {
	//data := []core.User{{0, "mock@gmail.ru", "Mock", "mock", true, false}}
	data := []core.User{{uuid.UUID{}, "mock@gmail.ru", "Mock", "mock", true, false}}
	return &UserMockRepository{users: data}
}

//func (rep *UserMockRepository) Create(gmail string) (int, error) {
//	user := core.User{Id: len(rep.users), Gmail: gmail, Nickname: "NewUser", Username: "mock" + strconv.Itoa(len(rep.users)), IsRegistered: false, HasProfilePic: false}
//	rep.users = append(rep.users, user)
//	return user.Id, nil
//}

func (rep *UserMockRepository) Exists(gmail string) bool {
	ind := findGmail(rep.users, gmail)
	return ind < len(rep.users)
}

func (rep *UserMockRepository) GetById(userId int) (core.User, error) {
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

func (rep *UserMockRepository) GetByGmail(gmail string) (core.User, error) {
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

func (rep *UserMockRepository) Change(userId int, u core.User) core.User {
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
