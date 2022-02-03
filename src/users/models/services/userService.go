package services

import (
	"errors"
	"go_project/src/users/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {

}

func (s *UserService) LoginUser(p *models.User) error {
	var err, isUserExist = userRepository.IsUserExistByEmail(p)
	if err != nil {
		return err
	}
	if isUserExist {
		var err, password = userRepository.GetUserPassword(p)

		if err != nil {
			return err
		}

		byteHash := []byte(password)
		bytePass := []byte(p.Password)
		result := bcrypt.CompareHashAndPassword(byteHash, bytePass)
		if result != nil {
			return errors.New("Login or password is not correct")
		}
		return userRepository.GetUsernameAndEmail(p)
	}

	return errors.New("Login or password is not correct")
}

func (s *UserService) UserRegister(p *models.User) error {
	var err, isUserExist = userRepository.IsUserExistByEmail(p)
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

	err = userRepository.CreateUser(p)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUsers (start, count int) ([]models.User, error) {
	return userRepository.GetUsers(start, count)
}

func (s *UserService) DeleteUser(p *models.User) error {
	return userRepository.DeleteUser(p)
}

func (s *UserService) UpdateUser(p *models.User) error {
	return userRepository.UpdateUser(p)
}

func (s *UserService) GetUser(p *models.User) error {
	return userRepository.GetUser(p)
}