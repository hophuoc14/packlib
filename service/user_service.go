package service

import (
	"packlib/model"
)

type UserService interface {
	Register(request model.CreateUserRequest) (model.CreateUserResponse, error)
	
	Login(request model.LoginRequest) (model.LoginSuccessResponse, error)
	
	FindUserByUsername(username string) (model.GetUserResponse, error)
}