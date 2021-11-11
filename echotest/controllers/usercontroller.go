package controllers

import (
	"fmt"
	"models"
)

func GetUser() {

}

func addUser(i int, u ...Usermodel) {

	var newuser = make(map[int]interface{})
	newuser[i] = u
	fmt.Println(newuser)
}

func deleteUser() {

}

func search(){
	func searchId(){}
	func searchName(){}
}