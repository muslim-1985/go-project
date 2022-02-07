package services

import (
	"go_project/internal/users/user/models"
)

type AuthServiceInterface interface {
	LoginUser(p *models.User) error
	UserRegister(p *models.User) error
}
