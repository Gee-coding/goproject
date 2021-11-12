package controllers

import (
	"fmt"
)

//models.Usermodel
func GetUser() {

}

func addUser(i int, u ...Usermodel) {

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
