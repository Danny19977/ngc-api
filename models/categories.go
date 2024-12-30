package model

import (
	"image"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model

	// id int
	Name string
	Image image.Image
	
}