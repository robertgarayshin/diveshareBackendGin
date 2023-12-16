package service

import (
	"diveshareBackendGin/model"
	"diveshareBackendGin/pkg/repository"
)

type RentService struct {
	repo repository.Rent
}

func (r RentService) Create(rent model.Rent) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r RentService) GetAll() ([]model.Rent, error) {
	//TODO implement me
	panic("implement me")
}

func (r RentService) GetById(rentId int) (model.Rent, error) {
	//TODO implement me
	panic("implement me")
}

func (r RentService) Update(rentId int, input model.UpdateRentInput) error {
	//TODO implement me
	panic("implement me")
}

func (r RentService) Delete(rentId int) error {
	//TODO implement me
	panic("implement me")
}

func NewRentService(repo repository.Rent) *RentService {
	return &RentService{repo: repo}
}
