package repository

import "gorm.io/gorm"

type BaseRepository struct {
	database *gorm.DB
}
