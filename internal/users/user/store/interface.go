package store

import (
	"go_project/internal/users/user/models"
)

type UserRepositoryInterface interface {
	GetUser(e *models.User) error
	UpdateUser(e *models.User) error
	DeleteUser(e *models.User) error
	IsUserExistByEmail(e *models.User) (error, bool)
	GetUserPassword(e *models.User) (error, string)
	GetUsernameAndEmail (e *models.User) error
	CreateUser(e *models.User) error
	GetUsers (start, count int) ([]models.User, error)
}
