package repository

import "packlib/entity"

type DepartmentFindParams struct {
	BaseQueryParams
	Name *string
}

type DepartmentRepository interface {
	Insert(department entity.Department) (entity.Department,error)
	Update(department entity.Department) (entity.Department,error)
	Delete(department entity.Department) error
	Find(param DepartmentFindParams) (*entity.Department, error)
}