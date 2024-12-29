package model

import "gorm.io/gorm"

type cart struct {
	gorm.Model

	user_id   int
	course_id int
}