package main

import (
	r "github.com/Valweb/AulasGolangWeb/routers"
	"github.com/labstack/echo/middleware"
	_ "github.com/labstack/echo/middleware"
)

//Primeira forma de fazer  um callback
func main() {
	e := r.App

	e.Use(middleware.Logger()) //Cada requisição ele estará exibindo no formato de log no terminal.
	e.Logger.Fatal(e.Start(":3000"))
}
