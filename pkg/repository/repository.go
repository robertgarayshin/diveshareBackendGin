package repository

import (
	"diveshareBackendGin/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username string) (model.User, error)
}

type User interface {
	GetAll() ([]model.User, error)
	GetById(userId int) (model.User, error)
	Delete(userId int) error
	Update(userId int, input model.User) error
}

//type TodoItem interface {
//	Create(listId int, item todo.TodoItem) (int, error)
//	GetAll(userId, listId int) ([]todo.TodoItem, error)
//	GetById(userId, itemId int) (todo.TodoItem, error)
//	Delete(userId, itemId int) error
//	Update(userId, itemId int, input todo.UpdateItemInput) error
//}

type Repository struct {
	Authorization
	User
	//TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		//TodoItem:      NewTodoItemPostgres(db),
	}
}
