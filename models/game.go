package models

//DB model
type Game struct {
	Name       string
	Plataform  string
	Status     string
	Percentage string
}

//Request
type GetGameRequest struct {
	Name      string
	Plataform string
}

type AddGameRequest struct {
	Name       string
	Plataform  string
	Status     string
	Percentage string
}

type RemoveGameRequest struct {
	Name      string
	Plataform string
}

//Response
type GetGameResponse struct {
	Name       string
	Plataform  string
	Status     string
	Percentage string
}

type AddGameResponse struct {
	Message string
}

type RemoveGameResponse struct {
	Message string
}
