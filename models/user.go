package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	// id                  int
	Name                string
	Postnom             string
	Prenom              string
	Tranche_age         string
	Email               string
	Telephone           string
	Bio                 string
	Role                string
	Permission          string
	Type_compte         string
	Status_compte       string
	Agree_terme_service bool
}
