package model

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Age      int8   `json:"age"`
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