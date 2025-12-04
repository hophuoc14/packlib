package service

import (
	"packlib/entity"
	"packlib/exception"
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
	user, err := service.userRepository.Login(request.Username, request.Password)
	exception.PanicNeeded(err)

	// token, er = 
}

func (service *UserServiceImpl) FindUserByUsername(username string) (model.GetUserResponse, error) {
	
}

