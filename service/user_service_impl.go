package service

import (
	"packlib/model"
	"packlib/repository"
	"packlib/validation"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (service *UserServiceImpl) Register(request model.CreateUserRequest) (model.CreateUserResponse, error) {
	validateUserErr := validation.ValidateUser(request)

	if validateUserErr != nil {
		return model.CreateUserResponse{}, validateUserErr
	}

}

func (service *UserServiceImpl) Login(request model.CreateUserRequest) (model.LoginSuccessResponse, error) {
	
}

func (service *UserServiceImpl) FindUserByUsername(username string) (model.GetUserResponse, error) {
	
}

