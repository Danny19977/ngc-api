package model

import "gorm.io/gorm"

type newslatter struct {
	gorm.Model
	email string
}