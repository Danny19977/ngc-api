package model

import(
"gorm.io/gorm"
) 



type courses struct{
	gorm.Model

	name string
	level int16
	category int32
	price  string
	status_courses bool
	is_valid  bool 
}