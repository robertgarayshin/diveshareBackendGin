package service

import (
	"diveshareBackendGin/model"
	"diveshareBackendGin/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(userId int) (model.User, error) {
	return s.repo.GetById(userId)
}

func (s *UserService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

func (s *UserService) Update(userId int, input model.UpdateUserInput) error {
	return s.repo.Update(userId, input)
}
