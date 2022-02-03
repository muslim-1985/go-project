package services

import "go_project/src/users/models"

type UserServiceInterface interface {
	LoginUser(p *models.User) error
	UserRegister(p *models.User) error
	GetUsers (start, count int) ([]models.User, error)
	DeleteUser(p *models.User) error
	UpdateUser(p *models.User) error
	GetUser(p *models.User) error
}
