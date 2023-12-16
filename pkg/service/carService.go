package service

import (
	"diveshareBackendGin/model"
	"diveshareBackendGin/pkg/repository"
)

type CarService struct {
	repo repository.Car
}

func (c CarService) Create(car model.Car) (int, error) {
	return c.repo.Create(car)
}

func (c CarService) GetAll() ([]model.Car, error) {
	return c.repo.GetAll()
}

func (c CarService) GetById(carId int) (model.Car, error) {
	return c.repo.GetById(carId)
}

func (c CarService) Update(carId int, input model.UpdateCarInput) error {
	return c.repo.Update(carId, input)
}

func (c CarService) Delete(carId int) error {
	return c.repo.Delete(carId)
}

func NewCarService(repo repository.Car) *CarService {
	return &CarService{repo: repo}
}
