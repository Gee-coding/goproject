package controllers

import (
	"goproject/echotest/models"
	"net/http"

	"github.com/labstack/echo"
)


var (
	users = map[int]*models.User{}
	seq   = 1
)

func getUserModel(c echo.Context) error {
	var m models.User

	return c.JSON(http.StatusOK, m)
}
func createNewUser()*User{
	return &User{}
}

func editUser(c echo.Context) error{

}
// func getUser(c echo.Context) error {
// 	u := []models.UserModel{
// 		{Id: 123,
// 			Name:     "John",
// 			Position: "addmin",
// 			Age:      20,
// 		},
// 		{Id: 234,
// 			Name:     "Peter",
// 			Position: "account",
// 			Age:      31,
// 		},
// 	} return

// }
