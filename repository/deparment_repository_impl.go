package repository

import (
	"packlib/entity"

	"gorm.io/gorm"
)

type DepartmentRepositoryImpl struct {
	BaseRepository
}

func NewDepartmentRepository(database *gorm.DB) DepartmentRepository {
	return &DepartmentRepositoryImpl{BaseRepository: BaseRepository{database: database}}
}

func (repository *DepartmentRepositoryImpl) Insert(param entity.Department) (entity.Department, error) {
	department := repository.database.Create(&param)

	if department.Error != nil {
		return entity.Department{}, department.Error
	}

	return param, nil
}

func (repository *DepartmentRepositoryImpl) Update(param entity.Department) (*entity.Department, error) {
	var updatedDepartment entity.Department

	repository.database.Save(&param).First(&updatedDepartment)
	return &updatedDepartment, nil
}

func (repository *DepartmentRepositoryImpl) Delete(param entity.Department) error {
	return repository.database.Delete(&param).Error
}

func (repository *DepartmentRepositoryImpl) Find(param DepartmentFindParams) ([]entity.Department, error) {
	var departments []entity.Department
	
	query := repository.database.Preload("Employees")
	
	// Only filter by name if provided
	if param.Name != nil && *param.Name != "" {
		query = query.Where("name = ?", *param.Name)
	}
	
	return departments, query.Find(&departments).Error
}
