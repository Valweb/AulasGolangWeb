package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

//Home é a página inicial da minha aplicação
func Home(c echo.Context) error {

	return c.Stringig(http.StatusOK, "Hello, World! GO GO")
}
