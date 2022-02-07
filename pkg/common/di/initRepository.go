package di

import (
	store2 "go_project/internal/users/user/store"
)

type AppRepository struct {
	userRepository store2.UserRepositoryInterface
}

func (a *AppRepository) InitRepository() *AppRepository {
	return &AppRepository{
		userRepository: &store2.UserRepository{},
	}
}
