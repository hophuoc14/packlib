package service

import (
	"packlib/entity"
	"packlib/exception"
	"packlib/model"
	"packlib/repository"
	"packlib/validation"

	"github.com/google/uuid"
)

type DepartmentServiceImpl struct {
	repository repository.DepartmentRepository
}

func NewDepartmentService(repository repository.DepartmentRepository) DepartmentService {
	return &DepartmentServiceImpl{repository: repository}
}

func (service *DepartmentServiceImpl) Insert(param entity.Department) (entity.Department, error) {
	error := validation.ValidateDepartment(param, true)
	if error != nil {
		return entity.Department{}, error
	}
	return service.repository.Insert(param)
}

func (service *DepartmentServiceImpl) Update(param entity.Department) (*entity.Department, error) {
	err := validation.ValidateDepartment(param, false)
	if err != nil {
		return nil, err
	}
	department, err := service.FindById(param.Id)
	if err != nil {
		return nil, err
	}
	if department == nil {
		return nil, exception.ValidationError{
			Message: "department not found",
		}	
	}
	return service.repository.Update(param)
}

func (service *DepartmentServiceImpl) Delete(param entity.Department) error {
	department, err := service.FindById(param.Id)
	if err != nil {
		return err
	}
	if department == nil {
		return exception.ValidationError{
			Message: "department not found",
		}	
	}
	return service.repository.Delete(param)
}

func (service *DepartmentServiceImpl) Find(param repository.DepartmentFindParams) ([]entity.Department, error) {
	return service.repository.Find(param)
}

func (service *DepartmentServiceImpl) FindById(id uuid.UUID) (*entity.Department, error) {
	limit := 1
	departments, err := service.repository.Find(repository.DepartmentFindParams{
		BaseQueryParams: model.BaseQueryParams{
			Limit: &limit,
			Id:    &id,
		},
	})
	if err != nil {
		return nil, err
	}
	if len(departments) < 1 {
		return nil, nil
	}
	return &departments[0], nil
}
