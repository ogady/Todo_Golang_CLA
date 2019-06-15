package infra

import (
	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo) {
	todoFormatter := InjectTodoFormatter()
	e.GET("/", todoFormatter.View())

	e.GET("/search", todoFormatter.Search())

	e.POST("/todoCreate", todoFormatter.Add())

	e.POST("/todoEdit", todoFormatter.Edit())
}
