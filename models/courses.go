package model

import(
"gorm.io/gorm"
) 



type Courses struct{
	gorm.Model

	// id int
	Name string
	Level int16
	Category uint `gorm:"foreignKey:category_id" json:"category_id"`
	Price  string
	Status_courses bool
	Is_valid  bool 
}