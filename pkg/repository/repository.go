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
	Update(userId int, input model.UpdateUserInput) error
}

type Car interface {
	Create(car model.Car) (int, error)
	GetAll() ([]model.Car, error)
	GetById(carId int) (model.Car, error)
	Update(carId int, input model.UpdateCarInput) error
	Delete(carId int) error
}

type Rent interface {
	Create(rent model.Rent) (int, error)
	GetAll() ([]model.Rent, error)
	GetById(rentId int) (model.Rent, error)
	Update(rentId int, input model.UpdateRentInput) error
	Delete(rentId int) error
}

type Review interface {
	Create(rent model.Review) (int, error)
	GetAll() ([]model.Review, error)
	GetById(reviewId int) (model.Review, error)
	Update(reviewId int, input model.UpdateReviewInput) error
	Delete(reviewId int) error
}

type Repository struct {
	Authorization
	User
	Car
	Rent
	Review
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Car:           NewCarPostgres(db),
		Rent:          NewRentPostgres(db),
		Review:        NewReviewPostgres(db),
	}
}
