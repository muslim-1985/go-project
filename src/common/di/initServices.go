package di

import "go_project/src/users/models/services"

type AppService struct {
	UserService services.UserServiceInterface
	AuthService services.AuthServiceInterface
}

func (a *AppService) InitService() *AppService  {
	appRepository := AppRepository{}
	return &AppService{
		UserService: &services.UserService{
			UserRepository: appRepository.InitRepository().userRepository,
		},
		AuthService: &services.AuthService{
			UserRepository: appRepository.InitRepository().userRepository,
		},
	}
}