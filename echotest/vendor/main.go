package main

import (
	_ "goproject/echotest/controllers"
	"goproject/echotest/db"
	_ "goproject/echotest/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo" // set GOPATH module
	"github.com/labstack/echo/middleware"
)

var conn = db.ConnectDB()

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	// e.GET("/users", controllers.GetUser)

	// e.POST("/users", controllers.InsertUserDB(123, "John", "admin", 25))

	// e.DELETE("/users/:id", controllers.DeleteUser)

	// e.PUT("/users", controllers.EditUser)

	e.Logger.Fatal(e.Start(":25060"))
}
