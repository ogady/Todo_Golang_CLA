package main

import (
	"fmt"
	"todo/infra/controller"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("sever start")
	e := echo.New()
	controller.InitRouting(e)
	e.Logger.Fatal(e.Start(":8080"))
}
