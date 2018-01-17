package main

import (
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/labstack/echo"
)

/*
Primeira forma de afzer um callback
func main() {
	e := echo.New()
	//criar uma rota para fazer uma requisição para o provedor
	e.GET("/", func(c echo.Context) error {
		//varial c será tipo context
		return c.String(http.StatusOK, "Hello, World!")

	})
	//Número da porta que o Go ficará escutando
	e.Start(":3000")
}
*/

//Segunda forma de fazer um callback
func main() {
	e := echo.New()
	e.GET("/", home)
	e.Start(":3000")
}

func home(c echo.Context) error {
	//varial c será tipo context
	return c.String(http.StatusOK, "Hello, World! GO GO")
}
