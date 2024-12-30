package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	// id int
	Number_oder int32
	CourseID    uint `gorm:"foreignKey:course_id" json:"course_id"`
	Status      string
}
