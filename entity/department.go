package entity

type Department struct {
	Entity
	Name        string     `gorm:"type:varchar(100);not null;unique;index" json:"name"`
	Description string     `gorm:"type:text" json:"description"`
	Employees   []Employee `gorm:"many2many:employee_departments;" json:"employees"`
}