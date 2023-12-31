package main

import (
	"net/http"

	"go-echo-template/internal/app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", routes.GetUser)

	e.Logger.Fatal(e.Start(":8080"))
}
