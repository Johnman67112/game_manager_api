package models

import "gorm.io/gorm"

//DB model
type Game struct {
	gorm.Model
	Name       string `json:"name"`
	Plataform  string `json:"plataform"`
	Status     string `json:"status"`
	Percentage string `json:"percentage"`
}
