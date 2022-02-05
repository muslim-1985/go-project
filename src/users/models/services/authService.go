package services

import (
	"errors"
	"go_project/src/users/models"
	"go_project/src/users/store"
)

type AuthService struct {
	UserRepository store.UserRepositoryInterface
}

func (s *AuthService) LoginUser(p *models.User) error {
	var err, isUserExist = s.UserRepository.IsUserExistByEmail(p)
	if err != nil {
		return err
	}
	if isUserExist {
		var err, password = s.UserRepository.GetUserPassword(p)

		if err != nil {
			return err
		}
		err = p.IsPasswordValid(password)

		if err != nil {
			return err
		}
		return s.UserRepository.GetUsernameAndEmail(p)
	}

	return errors.New("Login or password is not correct")
}

func (s *AuthService) UserRegister(p *models.User) error {
	var err, isUserExist = s.UserRepository.IsUserExistByEmail(p)
	if err != nil {
		return err
	}

	if isUserExist {
		return errors.New("A user is already registered to this mail")
	}

	err = p.CreatePasswordHash()

	if err != nil {
		return err
	}

	err = s.UserRepository.CreateUser(p)
	if err != nil {
		return err
	}
	return nil
}