package infra

import (
	"todo/domain/repository"

	"todo/interface/handler"
	"todo/usecase"
)

func InjectDB() SqlHandler {
	sqlhandler := NewSqlHandler()
	return *sqlhandler
}

func InjectTodoRepository() repository.TodoRepository {
	sqlHandler := InjectDB()
	return NewTodoRepository(sqlHandler)
}

func InjectTodoUsecase() usecase.TodoUsecase {
	TodoRepo := InjectTodoRepository()
	return usecase.NewTodoUsecase(TodoRepo)
}

func InjectTodoHandler() handler.TodoHandler {
	return handler.NewTodoHandler(InjectTodoUsecase())
}
