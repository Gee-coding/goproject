package models

type UserModel struct {
	Id       map[int]interface{}   `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Age      int    `json:"age"`
}
