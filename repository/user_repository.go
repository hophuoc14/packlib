package repository

import "packlib/entity"

type UserRepository interface {

	Insert(param entity.User) (entity.User, error)

	FindByUsername(username string) (entity.User, error)

	Login(username string, password string) (entity.User, error)

}