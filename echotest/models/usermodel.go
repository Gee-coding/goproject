package models

type User struct {
	Id       int  `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Salary      int    `json:"salary"`
}

