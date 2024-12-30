package model

import (
	// "golang.org/x/text/message"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	// id int
	Fullname string
	Email    string
	Subject  string
	Message  string
}
