package model

import "gorm.io/gorm"

type Paiement struct {
	gorm.Model

	// Id               int
	UserID           uint `gorm:"foreignKey:user_id" json:"user_id"`
	Methode_paiement string
	Montant          string
	CourseID         uint `gorm:"foreignKey:course_id" json:"course_id"`
}
