package infra

import (
	"todo/interface/handler"

	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo) {
	todoFormatter := new(handler.TodoFormatter)
	e.GET("/", todoFormatter.View())

	e.GET("/:word", todoFormatter.Search())

	e.POST("/todoCreate", todoFormatter.Add())

	e.POST("/todoEdit", todoFormatter.Edit())
}
