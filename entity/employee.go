package entity

import "github.com/google/uuid"

type Employee struct {
	Entity
	FullName string     `gorm:"type:varchar(255);not null" json:"full_name"`
	Phone    string     `gorm:"type:varchar(20);unique;index" json:"phone"`
	Age      int8       `gorm:"not null;type:smallint" json:"age"`
	Status   string     `gorm:"type:varchar(20);not null" json:"status"`
	UserId   *uuid.UUID `gorm:"type:varchar(50)" json:"user_id"`
}