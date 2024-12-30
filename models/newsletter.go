package model

import "gorm.io/gorm"

type Newslatter struct {
	gorm.Model

	Id int
	Email string
}