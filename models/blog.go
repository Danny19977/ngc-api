package model

import (
	"image"

	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model

	// id int
	Title       string
	Category_id int16
	Resume      string
	Content     int
	Image       image.Image
}
