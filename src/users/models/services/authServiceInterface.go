package services

import "go_project/src/users/models"

type AuthServiceInterface interface {
	LoginUser(p *models.User) error
	UserRegister(p *models.User) error
}
