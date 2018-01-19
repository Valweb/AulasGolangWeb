package main

import (
	r "github.com/Valweb/AulasGolangWeb/routers"
	_ "github.com/labstack/echo"
)

/*
Primeira forma de fazer  um callback
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
	r.App.Start(":3000")
}
