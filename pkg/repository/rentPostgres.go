package repository

import (
	"diveshareBackendGin/model"
	"github.com/jmoiron/sqlx"
)

type RentPostgres struct {
	db *sqlx.DB
}

func (r RentPostgres) Create(rent model.Rent) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r RentPostgres) GetAll() ([]model.Rent, error) {
	//TODO implement me
	panic("implement me")
}

func (r RentPostgres) GetById(rentId int) (model.Rent, error) {
	//TODO implement me
	panic("implement me")
}

func (r RentPostgres) Update(rentId int, input model.UpdateRentInput) error {
	//TODO implement me
	panic("implement me")
}

func (r RentPostgres) Delete(rentId int) error {
	//TODO implement me
	panic("implement me")
}

func NewRentPostgres(db *sqlx.DB) *RentPostgres {
	return &RentPostgres{db: db}
}
