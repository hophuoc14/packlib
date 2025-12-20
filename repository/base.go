package repository

import "gorm.io/gorm"

type BaseQueryParams struct {
	Id        *uint
	Offset    *int
	Limit     *int
	Sort      *string
	SortOrder *string
}

type BaseRepository struct {
	database *gorm.DB
}
