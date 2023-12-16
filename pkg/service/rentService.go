package service

import (
	"diveshareBackendGin/model"
	"diveshareBackendGin/pkg/repository"
)

type RentService struct {
	repo repository.Rent
}

func (r RentService) Create(rent model.Rent) (int, error) {
	return r.repo.Create(rent)
}

func (r RentService) GetAll() ([]model.Rent, error) {
	return r.repo.GetAll()
}

func (r RentService) GetById(rentId int) (model.Rent, error) {
	return r.repo.GetById(rentId)
}

func (r RentService) Update(rentId int, input model.UpdateRentInput) error {
	return r.repo.Update(rentId, input)
}

func (r RentService) Delete(rentId int) error {
	return r.repo.Delete(rentId)
}

func NewRentService(repo repository.Rent) *RentService {
	return &RentService{repo: repo}
}
