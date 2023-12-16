package repository

import (
	"diveshareBackendGin/model"
	"github.com/jmoiron/sqlx"
)

type CarPostgres struct {
	db *sqlx.DB
}

func (c CarPostgres) Create(car model.Car) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c CarPostgres) GetAll() ([]model.Car, error) {
	//TODO implement me
	panic("implement me")
}

func (c CarPostgres) GetById(carId int) (model.Car, error) {
	//TODO implement me
	panic("implement me")
}

func (c CarPostgres) Update(carId int, input model.UpdateCarInput) error {
	//TODO implement me
	panic("implement me")
}

func (c CarPostgres) Delete(carId int) error {
	//TODO implement me
	panic("implement me")
}

func NewCarPostgres(db *sqlx.DB) *CarPostgres {
	return &CarPostgres{db: db}
}
