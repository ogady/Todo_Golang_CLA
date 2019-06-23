package main

import (
	"fmt"
	"todo/infra/controller"
	"todo/injector"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("sever start")
	todoHandler := injector.InjectTodoHandler()
	e := echo.New()
	controller.InitRouting(e, todoHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
