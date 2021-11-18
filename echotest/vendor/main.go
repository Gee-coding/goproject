package main

import (
	"goproject/echotest/controllers"
	"goproject/echotest/db"
	"net/http"

	//"goproject/echotest/vendor/github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// set GOPATH module
)

var conn = db.ConnectDB()

func main() {
	e := echo.New()
	db.ConnectDB()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", func(c echo.Context) error {
		user := conn.QueryRow("")
		return c.JSON(http.StatusOK, user)
	})

	e.GET("/users/:user_id", controllers.GetUserDB)
	e.DELETE("/users/:user_id", controllers.DeleteUserDB)
	e.PUT("/users", controllers.EditUser)

	e.Logger.Fatal(e.Start(":25060"))
}
