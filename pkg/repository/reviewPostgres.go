package repository

import (
	"diveshareBackendGin/model"
	"github.com/jmoiron/sqlx"
)

type ReviewPostgres struct {
	db *sqlx.DB
}

func (r ReviewPostgres) Create(rent model.Review) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReviewPostgres) GetAll() ([]model.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReviewPostgres) GetById(reviewId int) (model.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReviewPostgres) Update(reviewId int, input model.UpdateReviewInput) error {
	//TODO implement me
	panic("implement me")
}

func (r ReviewPostgres) Delete(reviewId int) error {
	//TODO implement me
	panic("implement me")
}

func NewReviewPostgres(db *sqlx.DB) *ReviewPostgres {
	return &ReviewPostgres{db: db}
}
