package controllers

import (
	"fmt"
	"goproject/echotest/models"
)

//models.Usermodel
func GetUser() {

}

func addUser(i int, u ...models.UserModel) {

	var newuser = make(map[int]interface{})
	newuser[i] = u

	return
	fmt.Println(newuser)
}

func deleteUser() {

}

// func search(){
// 	func searchId(){}
// 	func searchName(){}
// }
