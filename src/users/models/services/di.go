package services

import "go_project/src/users/store"

var userRepository store.UserRepositoryInterface = &store.UserRepository{}
