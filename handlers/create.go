package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/minaduking/echo_sample/models"
)

func CreateUser(c echo.Context) (err error) {
	u := &models.User{
		ID : models.Seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	models.Users[u.ID] = u
	models.Seq++
	return c.JSON(http.StatusCreated, u)
}