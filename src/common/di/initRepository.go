package di

import "go_project/src/users/store"

type AppRepository struct {
	userRepository store.UserRepositoryInterface
}

func (a *AppRepository) InitRepository() *AppRepository  {
	return &AppRepository{
		userRepository: &store.UserRepository{},
	}
}
