package models

type User struct {
	Id       int  `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Age      int    `json:"age"`
}

type UserList struct{
	Id       int  `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Age      int    `json:"age"`
}
