package controllers

import (
	"database/sql"
	"goproject/echotest/models"
	"net/http"

	"github.com/labstack/echo"
)

// func GetUserModel(c echo.Context) error {
// 	var m models.User

// 	return c.JSON(http.StatusOK, m)
// }

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

func InsertUserDB(db *sql.DB, u models.User) error {
	stmt, err := db.Prepare("INSERT INTO user (user_id,user_name,user_position,user_age)VALUE(?,?,?,?)")
	_, err = stmt.Exec(u.Id, u.Name, u.Position, u.Age)
	return err
}

