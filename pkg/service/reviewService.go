package service

import (
	"diveshareBackendGin/model"
	"diveshareBackendGin/pkg/repository"
)

type ReviewService struct {
	repo repository.Review
}

func (r ReviewService) Create(rent model.Review) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReviewService) GetAll() ([]model.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReviewService) GetById(reviewId int) (model.Review, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReviewService) Update(reviewId int, input model.UpdateReviewInput) error {
	//TODO implement me
	panic("implement me")
}

func (r ReviewService) Delete(reviewId int) error {
	//TODO implement me
	panic("implement me")
}

func NewReviewService(repo repository.Review) *ReviewService {
	return &ReviewService{repo: repo}
}
