package model

import "gorm.io/gorm"

type order struct {
	gorm.Model

	number_oder int32
	course_id   int16
	status      string
}