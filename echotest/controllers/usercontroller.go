package controllers

import (
	"fmt"
	"goproject/echotest/db"
	"goproject/echotest/models"
	"net/http"

	"github.com/labstack/echo"
)

// func GetUserModel(c echo.Context) error {
// 	var m models.User

// 	return c.JSON(http.StatusOK, m)
// }
var conn = db.ConnectDB()

func InsertUser(c echo.Context) error {
	edit := new(models.User)
	if err := c.Bind(edit); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, edit)
}
func GetUser(c echo.Context) error {
	u := []models.User{
		{Id: 123,
			Name:     "John",
			Position: "addmin",
			Age:      20,
		},
		{Id: 234,
			Name:     "Peter",
			Position: "account",
			Age:      31,
		},
	}
	return c.JSON(200, u)

}

func DeleteUser(c echo.Context) error {

	id := c.Param("id")
	return c.String(http.StatusOK, id)

}
func EditUser(c echo.Context) error {
	id := c.QueryParam("id")
	return c.String(http.StatusOK, id)

}

// func InsertUserDB(c echo.Context) error {

// }

func DeleteUserDB(c echo.Context) error {

	requested_id := c.Param("user_id")
	sql := "Delete FROM user Where user_id = ?"
	stmt, err := conn.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	result, err2 := stmt.Exec(requested_id)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(result.RowsAffected())
	return c.JSON(http.StatusOK, "Deleted")
}
