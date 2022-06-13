package model

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Todos *[]Todo `json:"todos"`
}
