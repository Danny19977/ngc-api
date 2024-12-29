package model

import "gorm.io/gorm"

type whitelist struct {
	gorm.Model

	user_id   int
	course_id int
}