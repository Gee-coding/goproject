package main

import (
	"goproject/echotest/controllers"
	"goproject/echotest/db"
	"net/http"

	//"goproject/echotest/vendor/github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	// set GOPATH module
)

var conn = db.ConnectDB()

func main() {
	e := echo.New()
	db.ConnectDB()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", controllers.GetAllUser)
	e.GET("/users/:user_id", controllers.GetUser)
	e.DELETE("/users/:user_id", controllers.DeleteUser)
	e.PUT("/users/:user_id/:user_name/:user_position/:user_salary", controllers.EditUser)
	e.POST("/users/:user_name/:user_position/:user_salary", controllers.AddUser)

	e.Logger.Fatal(e.Start(":25060"))
}
