package di

import (
	services2 "go_project/internal/users/user/models/services"
)

type AppService struct {
	UserService services2.UserServiceInterface
	AuthService services2.AuthServiceInterface
}

func (a *AppService) InitService() *AppService {
	appRepository := AppRepository{}
	return &AppService{
		UserService: &services2.UserService{
			UserRepository: appRepository.InitRepository().userRepository,
		},
		AuthService: &services2.AuthService{
			UserRepository: appRepository.InitRepository().userRepository,
		},
	}
}