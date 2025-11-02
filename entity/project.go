package entity

type Project struct {
	Entity
	Name         string     `gorm:"type:varchar(100);not null;unique;index" json:"name"`
	Description  string     `gorm:"type:text" json:"description"`
	DepartmentID string     `gorm:"type:varchar(50);index" json:"department_id"`
	Department   Department `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"department"`
}