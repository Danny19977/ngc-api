package model

import (
	"image"

	"gorm.io/gorm"
)

type category struct {
	gorm.Model

	name string
	image image.Image
	
}