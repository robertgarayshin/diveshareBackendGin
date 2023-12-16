package service

import (
	"diveshareBackendGin/model"
	"diveshareBackendGin/pkg/repository"
)

type CarService struct {
	repo repository.Car
}

func (c CarService) Create(car model.Car) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c CarService) GetAll() ([]model.Car, error) {
	//TODO implement me
	panic("implement me")
}

func (c CarService) GetById(carId int) (model.Car, error) {
	//TODO implement me
	panic("implement me")
}

func (c CarService) Update(carId int, input model.UpdateCarInput) error {
	//TODO implement me
	panic("implement me")
}

func (c CarService) Delete(carId int) error {
	//TODO implement me
	panic("implement me")
}

func NewCarService(repo repository.Car) *CarService {
	return &CarService{repo: repo}
}
