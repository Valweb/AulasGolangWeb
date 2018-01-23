package routers

import (
	"github.com/Valweb/AulasGolangWeb/controllers"
	"github.com/labstack/echo"
)

//App é uma instância de echo
var App *echo.Echo

func init() {
	App = echo.New()

	//Nossas Rotas
	App.GET("/", controllers.Home)
}
