package models

import "gopkg.in/validator.v2"

//DB model
type Game struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name" validate:"nonzero"`
	Plataform  string `json:"plataform" validate:"nonzero"`
	Status     string `json:"status" validate:"min=1, max=50"`
	Percentage string `json:"percentage" validate:"min=2, max=4"`
}

func GameValidator(game *Game) error {
	if err := validator.Validate(game); err != nil {
		return err
	}

	return nil
}
