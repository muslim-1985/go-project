package di

import "go_project/src/users/models/services"

type AppService struct {
	UserService services.UserServiceInterface
}

func (a *AppService) InitService() *AppService  {
	appRepository := AppRepository{}
	return &AppService{
		UserService: &services.UserService{
			UserRepository: appRepository.InitRepository().userRepository,
		},
	}
}