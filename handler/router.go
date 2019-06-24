package handler

import (
	"github.com/labstack/echo"
)

func InitRouting(e *echo.Echo, todoHandler TodoHandler) {

	e.GET("/", todoHandler.View())

	e.GET("/search", todoHandler.Search())

	e.POST("/todoCreate", todoHandler.Add())

	e.POST("/todoEdit", todoHandler.Edit())
}
