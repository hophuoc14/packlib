package entity

type User struct {
	Entity
	Username   string   `gorm:"type:varchar(100);unique;index" json:"username"`
	Password   string   `gorm:"type:varchar(100);not null" json:"-"`
	Email      string   `gorm:"type:varchar(100);unique;index" json:"email"`
	EmployeeID string   `gorm:"type:varchar(50);unique;index" json:"employee_id"`
	Employee   Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"employee"`
}