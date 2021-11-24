package controllers

import (
	"fmt"
	"goproject/echotest/db"
	"goproject/echotest/models"
	"net/http"
	"strconv"
	"github.com/labstack/echo"
)

var conn = db.ConnectDB()



func AddUser(c echo.Context) error {
	name := c.Param("user_name")
	position := c.Param("user_position")
	salary := c.Param("user_salary")

	stmt, err := conn.Prepare("INSERT INTO user (user_name,user_position,user_salary) VALUES (?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec(name, position, salary)
	if err != nil {
		fmt.Println(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, lastId)
}

func DeleteUser(c echo.Context) error {

	id := c.Param("user_id")
	stm := "DELETE FROM user WHERE user_id = ?"
	err, _ := conn.Exec(stm, id)
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, id)
}


func GetUser(c echo.Context) error {
	var (
		userModel     models.User
		userModelList []models.User
	)
	id := c.Param("user_id")
	err := conn.QueryRow(`SELECT * FROM user WHERE user_id = ?`, id).Scan(
		&userModel.Id, &userModel.Name, &userModel.Position, &userModel.Salary,
	)
	if err != nil {
		fmt.Println(err)
	}
	userModelList = append(userModelList, userModel)
	return c.JSON(http.StatusOK, userModelList)
}

func GetAllUser(c echo.Context) error {
	var (
		userModel     models.User
		userModelList []models.User
	)
	rows, err := conn.Query("SELECT * FROM user")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&userModel.Id, &userModel.Name, &userModel.Position, &userModel.Salary); err != nil {
			return err
		}
		userModelList = append(userModelList, userModel)
	}
	if err = rows.Err(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, userModelList)
}

func EditUser (c echo.Context) error {
	
	var (
		users = map[int]*models.User{}
	)
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("user_id"))
	users[id].Name = c.Param("user_name")
	return c.JSON(http.StatusOK, users[id])
}