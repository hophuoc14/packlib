package service

import (
	"packlib/entity"
	"packlib/repository"
)

type DepartmentService interface {
	Insert(param entity.Department) (entity.Department, error)
	Update(param entity.Department) (entity.Department, error)
	Delete(param entity.Department) error
	Find(param repository.DepartmentFindParams) (*entity.Department, error)
}