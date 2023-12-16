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
	Create(review model.Review) (int, error)
	GetAll() ([]model.Review, error)
	GetById(reviewId int) (model.Review, error)
	Update(reviewId int, input model.UpdateReviewInput) error
	Delete(reviewId int) error
}

type Service struct {
	Authorization
	User
	Car
	Rent
	Review
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		Car:           NewCarService(repos.Car),
		Rent:          NewRentService(repos.Rent),
		Review:        NewReviewService(repos.Review),
	}
}
