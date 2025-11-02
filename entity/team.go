package entity

type Team struct {
	Entity
	Name       string     `gorm:"type:varchar(255);not null;unique;index" json:"name"`
	EmployeeID string     `gorm:"type:varchar(50);index" json:"employee_id"`
	Employee   Employee   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"employee"`
	ProjectID  string     `gorm:"type:varchar(50);index" json:"project_id"`
	Project    Project    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"project"`
	Employees  []Employee `gorm:"many2many:employee_works;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"employees"`
}