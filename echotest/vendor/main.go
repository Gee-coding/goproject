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
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", func(c echo.Context) error {
		user := conn.QueryRow("")
		return c.String(http.StatusOK, user)
	})

	//e.POST("/users", controllers.InsertUserDB(123, "John", "admin", 25))

	e.DELETE("/users/:id", controllers.DeleteUser)

	e.PUT("/users", controllers.EditUser)

	// sql := "INSERT INTO test_ball.user(user_id, user_name, user_position,user_age) VALUES ('', 'Peter','sale','30')"

	// _, err := conn.Query(sql)

	// if err != nil {
	// 	panic(err.Error())
	// }

	e.Logger.Fatal(e.Start(":25060"))
}
