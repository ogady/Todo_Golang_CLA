package handler

import (
	"net/http"
	"todo/domain/model"
	"todo/usecase"

	"github.com/labstack/echo"
)

type TodoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) TodoHandler {
	todoHandler := TodoHandler{todoUsecase: todoUsecase}
	return todoHandler
}

func (handler *TodoHandler) View() echo.HandlerFunc {

	return func(c echo.Context) error {
		models, err := handler.todoUsecase.View()
		if err != nil {
			return c.JSON(http.StatusBadRequest, models)
		}
		return c.JSON(http.StatusOK, models)
	}

}
func (handler *TodoHandler) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		word := c.QueryParam("word")
		models, err := handler.todoUsecase.Search(word)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models)
		}
		return c.JSON(http.StatusOK, models)
	}
}
func (handler *TodoHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo model.Todo
		c.Bind(&todo)
		err := handler.todoUsecase.Add(&todo)
		return c.JSON(http.StatusOK, err)
	}
}
func (handler *TodoHandler) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo model.Todo
		c.Bind(&todo)
		err := handler.todoUsecase.Edit(&todo)
		return c.JSON(http.StatusOK, err)
	}
}
