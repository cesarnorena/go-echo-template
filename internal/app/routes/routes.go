package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
