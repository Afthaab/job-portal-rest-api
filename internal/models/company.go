package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	Field    string `json:"field" validate:"required"`
}
