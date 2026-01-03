package repository

import (
	"packlib/entity"
	"packlib/exception"
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
		return entity.User{}, exception.ValidationError{
			Message: "Username already exists",
		}
	}

	// Create user with employee association
	result := database.Session(&gorm.Session{FullSaveAssociations: true}).Create(&param)

	if result.Error != nil {
		return entity.User{}, result.Error
	}

	// Reload the user with employee to get the generated UUIDs
	var createdUser entity.User
	database.Preload("Employee").First(&createdUser, "id = ?", param.Id)

	return createdUser, nil

}

func (repository UserRepositoryImpl) FindByUsername(username string) (entity.User, error) {
	
	database := repository.database

	var user entity.User
	database.Where("username = ?", username).Preload("Employee").First(&user)
	
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
