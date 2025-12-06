package model

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required,min=2,max=100"`
	Phone    string `json:"phone" validate:"required,min=10,max=15"`
	Age      int8   `json:"age" validate:"required,min=1,max=150"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserResponse struct {
	Id uuid.UUID `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Employee CreatEmployeeResponse `json:"employee"`
}

type CreatEmployeeResponse struct {
	Id uuid.UUID `json:"id"`
	FullName string `json:"full_name"`
	Phone string `json:"phone"`
	Age int8 `json:"age"`
}

type GetUserResponse struct {
	Id uuid.UUID `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Employee CreatEmployeeResponse `json:"employee"`
}

type LoginSuccessResponse struct {
	Token Token `json:"token"`
	User GetUserResponse `json:"user"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}