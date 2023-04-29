package repository

import "server/internal/core"

type Authorization interface {
	SingIn(user core.User) (int, error)
}

type ReleaseItem interface {
	GetById(releaseId int) (core.Song, error)
}

type Repository struct {
	//Authorization
	ReleaseItem
}

func New(a int) *Repository {
	return &Repository{
		ReleaseItem: NewReleaseRepo(),
		//Authorization: ,
	}
}
