package main

import (
	"net/http"
	_"goproject/echotest/models"

	"github.com/labstack/echo" // set GOPATH module
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})


	e.POST("/users", func(c echo.Context) (err error) {
		u := new(models.UserModel)
		if err = c.Bind(u); err != nil{
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK,u)
	})
	e.GET("/users", getUserModel)
	e.Logger.Fatal(e.Start(":8080"))
}
