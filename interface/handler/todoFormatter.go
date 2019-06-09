package handler

import (
	"net/http"
	"todo/domain/model"
	"todo/usecase"

	"github.com/labstack/echo"
)

type TodoFormatter struct {
	todoUsecase usecase.TodoUsecase
}

func (formatter *TodoFormatter) View() echo.HandlerFunc {

	return func(c echo.Context) error {
		models, err := formatter.todoUsecase.View()
		if err != nil {
			return c.JSON(http.StatusBadRequest, models)
		}
		return c.JSON(http.StatusOK, models)
	}

}
func (formatter *TodoFormatter) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		word := c.QueryParam("word")
		models, err := formatter.todoUsecase.Search(word)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models)
		}
		return c.JSON(http.StatusOK, models)
	}
}
func (formatter *TodoFormatter) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo model.Todo
		c.Bind(&todo)
		err := formatter.todoUsecase.Add(&todo)
		return c.JSON(http.StatusOK, err)
	}
}
func (formatter *TodoFormatter) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo model.Todo
		c.Bind(&todo)
		err := formatter.todoUsecase.Edit(&todo)
		return c.JSON(http.StatusOK, err)
	}
}
