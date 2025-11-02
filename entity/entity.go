package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entity struct {
	Id uuid.UUID `gorm:"type:varchar(50);primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (e *Entity) BeforeCreate(db *gorm.DB) (err error) {
	e.Id = uuid.New()
	return
}