package model

import (
    "gorm.io/gorm"
)

type Support struct {
    gorm.Model
    ID          uint   `gorm:"primaryKey"`
    Fullname    string
    Email       string
    Telephone   string
    Subject     string
    Description string
}
