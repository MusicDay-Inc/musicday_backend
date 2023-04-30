package service

import (
	"server/internal/core"
	"server/internal/repository"
)

type Token interface {
	GetJWT(gmail string) (core.JWT, error)
	ParseToken(token string) (int, error)
}

type User interface {
	RegisterUser(userId int, user core.User) (core.User, error)
}

type Service struct {
	Token
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Token: NewTokenService(repos.User),
		User:  NewUserService(repos.User),
	}
}

//type List interface {
//	Create(userId int, list todo.TodoList) (int, error)
//	GetAll(userId int) ([]todo.TodoList, error)
//	GetById(userId, listId int) (todo.TodoList, error)
//	Delete(userId, listId int) error
//	Update(userId, listId int, input todo.UpdateListInput) error
//}
//
//type Item interface {
//	Create(userId, listId int, item todo.TodoItem) (int, error)
//	GetAll(userId, listId int) ([]todo.TodoItem, error)
//	GetById(userId, itemId int) (todo.TodoItem, error)
//	Delete(userId, itemId int) error
//	Update(userId, itemId int, input todo.UpdateItemInput) error
//}
