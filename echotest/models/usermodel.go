package models

type UserModel struct {
	Id       int  `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Age      int    `json:"age"`
}
