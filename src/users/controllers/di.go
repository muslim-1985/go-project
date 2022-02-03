package controllers

import "go_project/src/users/models/services"

var userService services.UserServiceInterface = &services.UserService{}