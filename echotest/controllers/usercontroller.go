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

func DeleteUser(c echo.Context) error {

	id := c.Param("id")
	return c.String(http.StatusOK, id)

}
func EditUser(c echo.Context) error {
	id := c.QueryParam("id")
	return c.String(http.StatusOK, id)

}

// func GetUser(c echo.Context) error {
// 	var (
// 		userModel models.User
// 		userModelList []models.UserList
// 	)
// id := c.QueryParam("user_id")
// fmt.Println(id)
// rows, err := conn.Query(`SELECT * FROM user WHERE user_id = ? `,id)
// if err != nil {
// 	fmt.Println("error user_id :", err)
// }
// for rows.Next() {
// 	err = rows.Scan(
// 		&userModel.Id,
// 		&userModel.Name,
// 		&userModel.Position,
// 		&userModel.Age,
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		userModelList = append(userModelList,models.UserList(userModel))
// 	}
// }
// rows.Close()
// return c.JSON(http.StatusOK, echo.Map{
// 	"status": true,
// 	"result": userModelList,
// })

// 	}

func GetUser(c echo.Context) error {
	var (
		userModel     models.User
		userModelList []models.User
	)
	id := c.QueryParam("user_id")

	rows, err := conn.Query(`SELECT * FROM user WHERE user_id = ?`, id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&userModel.Id, &userModel.Name, &userModel.Position, &userModel.Age)
	}
	if err != nil {
		fmt.Println(err)
	}
	userModelList = append(userModelList, userModel)
	rows.Close()
	return c.JSON(http.StatusOK, userModelList)
}
