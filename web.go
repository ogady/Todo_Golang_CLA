package main

import (
	"fmt"
	"todo/infra"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("sever start")
	e := echo.New()
	infra.InitRouting(e)
	e.Logger.Fatal(e.Start(":8080"))
}
