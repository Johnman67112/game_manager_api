package models

//DB model
type Game struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name"`
	Plataform  string `json:"plataform"`
	Status     string `json:"status"`
	Percentage string `json:"percentage"`
}
