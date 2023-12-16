package service

import (
	"diveshareBackendGin/model"
	"diveshareBackendGin/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	GetAll() ([]model.User, error)
	GetById(userId int) (model.User, error)
	Delete(userId int) error
	Update(userId int, input model.User) error
}

//
//type TodoItem interface {
//	Create(userId, listId int, item todo.TodoItem) (int, error)
//	GetAll(userId, listId int) ([]todo.TodoItem, error)
//	GetById(userId, itemId int) (todo.TodoItem, error)
//	Delete(userId, itemId int) error
//	Update(userId, itemId int, input todo.UpdateItemInput) error
//}

type Service struct {
	Authorization
	User
	//TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		//TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
