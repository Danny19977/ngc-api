package model

import "gorm.io/gorm"

type user struct {
	gorm.Model

	name        string
	postnom     string
	prenom      string
	tranche_age string
	email       string
	telephone   string
	bio         string
	role        string
	permission  string
}