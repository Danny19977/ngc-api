package model

import "gorm.io/gorm"

type Whitelist struct {
	gorm.Model
	// id int
	UserID   uint `gorm:"foreignKey:user_id" json:"user_id"`
	CourseID uint `gorm:"foreignKey:course_id" json:"course_id"`
}
