package entity

type User struct {
	Entity
	Username string   `gorm:"type:varchar(100);unique;index" json:"username"`
	Password string   `gorm:"type:varchar(100);not null" json:"-"`
	Email    string   `gorm:"type:varchar(100);unique;index" json:"email"`
	Employee Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"employee"`
}