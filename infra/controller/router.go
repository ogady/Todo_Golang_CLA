package controller

import (
	"todo/interface/handler"

	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, todoHandler handler.TodoHandler) {

	e.GET("/", todoHandler.View())

	e.GET("/search", todoHandler.Search())

	e.POST("/todoCreate", todoHandler.Add())

	e.POST("/todoEdit", todoHandler.Edit())
}
