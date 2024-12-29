package model

import "gorm.io/gorm"

type paiement struct {
	gorm.Model
	user_id          int
	methode_paiement string
	montant          string
	course_id        int
}