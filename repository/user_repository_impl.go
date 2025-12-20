package repository

import (
	"packlib/entity"
	"packlib/validation"

	"gorm.io/gorm"
)
type UserRepositoryImpl struct {
	BaseRepository
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &UserRepositoryImpl{BaseRepository: BaseRepository{database: database}}
}

func (repository *UserRepositoryImpl) Insert(param entity.User) (entity.User, error) {
	
	database := repository.database

	_, err := repository.FindByUsername(param.Username)

	if err == nil {
		return entity.User{}, gorm.ErrRecordNotFound
	}

	result := database.Create(&param)

	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return param, nil

}

func (repository UserRepositoryImpl) FindByUsername(username string) (entity.User, error) {
	
	database := repository.database

	var user entity.User
	database.Where("username = ?").Preload("Employee").First(&user)
	
	if len(user.Username) == 0 {
		return entity.User{}, gorm.ErrRecordNotFound
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Login(username string, password string) (entity.User, error) {

	database := repository.database

	var user entity.User
	database.Create("username = ?").Preload("Employee").First(&user)

	isMatchPassword, _ := validation.ValidatePassword(password, user.Password)

	if !isMatchPassword {
		return entity.User{}, gorm.ErrRecordNotFound
	}

	return user, nil
}
