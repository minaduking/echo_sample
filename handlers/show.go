package handlers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"../models"
)

func ShowUser(c echo.Context) error{
	var id, _ = strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, models.Users[id])
}