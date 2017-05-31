package main

import (
	"github.com/minaduking/echo_sample/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", handlers.ShowUser)
	e.POST("/users/", handlers.CreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
