package repository

import (
	"server/internal/core"
)

type Authorization interface {
	SingIn(user core.User) (int, error)
}

type ReleaseItem interface {
	GetById(releaseId int) (core.Song, error)
}

type User interface {
	// Create returns id of new user, and changes his id
	Create(gmail string) (int, error)
	Exists(gmail string) bool
	GetById(userId int) (core.User, error)
	GetByGmail(gmail string) (core.User, error)
}

type Repository struct {
	//Authorization
	ReleaseItem
	User
}

func New(DB string) *Repository {
	return &Repository{
		ReleaseItem: NewReleaseRepo(),
		User:        NewUserRepo(),
		//Authorization: ,
	}
}
