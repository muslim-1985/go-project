package services

import (
	"go_project/src/users/models"
	"go_project/src/users/store"
)

type UserService struct {
	UserRepository store.UserRepositoryInterface
}

func (s *UserService) GetUsers (start, count int) ([]models.User, error) {
	return s.UserRepository.GetUsers(start, count)
}

func (s *UserService) DeleteUser(p *models.User) error {
	return s.UserRepository.DeleteUser(p)
}

func (s *UserService) UpdateUser(p *models.User) error {
	return s.UserRepository.UpdateUser(p)
}

func (s *UserService) GetUser(p *models.User) error {
	return s.UserRepository.GetUser(p)
}