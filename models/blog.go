package model

import (
	"image"

	"gorm.io/gorm"
)

type blog struct {
	gorm.Model

	title       string
	category_id int
	resume      string
	content     int
	image       image.Image
}
