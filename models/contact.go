package model

import (
	// "golang.org/x/text/message"
	"gorm.io/gorm"
)

type contact struct {
	gorm.Model

	fullname string
	email    string
	subject  string
	message string
}