package service

import (
	"diveshareBackendGin/model"
	"diveshareBackendGin/pkg/repository"
)

type ReviewService struct {
	repo repository.Review
}

func (r ReviewService) Create(review model.Review) (int, error) {
	return r.repo.Create(review)
}

func (r ReviewService) GetAll() ([]model.Review, error) {
	return r.repo.GetAll()
}

func (r ReviewService) GetById(reviewId int) (model.Review, error) {
	return r.repo.GetById(reviewId)
}

func (r ReviewService) Update(reviewId int, input model.UpdateReviewInput) error {
	return r.repo.Update(reviewId, input)
}

func (r ReviewService) Delete(reviewId int) error {
	return r.repo.Delete(reviewId)
}

func NewReviewService(repo repository.Review) *ReviewService {
	return &ReviewService{repo: repo}
}
