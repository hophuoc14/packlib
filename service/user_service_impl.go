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

	hashedPassword, err := validation.HashPassword(request.Password)
	exception.PanicNeeded(err)

	user, err := service.userRepository.Insert(entity.User{
		Username: request.Username,
		Password: hashedPassword,
		Email: request.Email,
		Employee: entity.Employee{
			FullName: request.FullName,
			Phone: request.Phone,
			Age: request.Age,
		},
	})
	exception.PanicNeeded(err)

	return model.CreateUserResponse{
		Username: user.Username,
		Email: user.Email,
		Employee: model.CreatEmployeeResponse{
			Id: user.Employee.Id,
			FullName: user.Employee.FullName,
			Phone: user.Employee.Phone,
			Age: user.Employee.Age,
		},
	}, nil
}

func (service *UserServiceImpl) Login(request model.CreateUserRequest) (model.LoginSuccessResponse, error) {
	user, err := service.userRepository.Login(request.Username, request.Password)
	exception.PanicNeeded(err)

	token, err := validation.CreateToken(user)
	exception.PanicNeeded(err)

	return model.LoginSuccessResponse{
		Token: model.Token{
			AccessToken: token,
		},
		User: model.GetUserResponse{
			Username: user.Username,
			Email: user.Email,
			Employee: model.CreatEmployeeResponse{
				Id: user.Employee.Id,
				FullName: user.Employee.FullName,
				Phone: user.Employee.Phone,
				Age: user.Employee.Age,
			},
		},
	}, nil
}

func (service *UserServiceImpl) FindUserByUsername(username string) (model.GetUserResponse, error) {
	user, err := service.userRepository.FindByUsername(username)
	exception.PanicNeeded(err)

	return model.GetUserResponse{
		Username: user.Username,
		Email: user.Email,
		Employee: model.CreatEmployeeResponse{
			Id: user.Employee.Id,
			FullName: user.Employee.FullName,
			Phone: user.Employee.Phone,
			Age: user.Employee.Age,
		},
	}, nil
}

