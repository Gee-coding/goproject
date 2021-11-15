package controllers

import (
	"goproject/echotest/models"
	"net/http"

	"github.com/labstack/echo"
)

var (
	users = map[int]*models.UserModel{}
	seq   = 1
)

func getUserModel(c echo.Context) error {
	var m models.UserModel

	return c.JSON(http.StatusOK, m)
}
func createUser(c echo.Context) error {
	u := &models.UserModel{
		Id : seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.Id] = u
	seq++
	return c.JSON(http.StatusCreated, u)
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
