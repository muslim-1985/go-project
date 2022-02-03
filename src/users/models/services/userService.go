package services

import (
	"errors"
	"go_project/src/users/models"
	"go_project/src/users/store"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository store.UserRepositoryInterface
}


func (s *UserService) LoginUser(p *models.User) error {
	var err, isUserExist = s.UserRepository.IsUserExistByEmail(p)
	if err != nil {
		return err
	}
	if isUserExist {
		var err, password = s.UserRepository.GetUserPassword(p)

		if err != nil {
			return err
		}

		byteHash := []byte(password)
		bytePass := []byte(p.Password)
		result := bcrypt.CompareHashAndPassword(byteHash, bytePass)
		if result != nil {
			return errors.New("Login or password is not correct")
		}
		return s.UserRepository.GetUsernameAndEmail(p)
	}

	return errors.New("Login or password is not correct")
}

func (s *UserService) UserRegister(p *models.User) error {
	var err, isUserExist = s.UserRepository.IsUserExistByEmail(p)
	if err != nil {
		return err
	}

	if isUserExist {
		return errors.New("A user is already registered to this mail")
	}

	bytePassword := []byte(p.Password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.MinCost)
	if err != nil {
		return err
	}
	password := string(hash)
	p.Password = password

	err = s.UserRepository.CreateUser(p)
	if err != nil {
		return err
	}
	return nil
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