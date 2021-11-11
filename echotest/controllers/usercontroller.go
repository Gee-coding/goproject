package controllers

import (
	"fmt"
	"github.com/goproject/models"
)

func GetUser() {

}

func addUser(i int, u ...usermodel.Usermodel) {

	var newuser = make(map[int]interface{})
	newuser[i] = u
	fmt.Println(newuser)
}

func deleteUser() {

}
