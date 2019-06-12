package infra

import (
	"todo/domain/repository"
	"todo/interface/database"
	"todo/interface/handler"
	"todo/usecase"
)

func InjectDB() database.SqlHandler {
	sqlhandler := NewSqlHandler()
	return database.SqlHandler(sqlhandler)
}

func InjectTodoRepository() repository.TodoRepository {
	sqlHandler := InjectDB()
	return NewTodoRepository(sqlHandler)
}

func InjectTodoUsecase() usecase.TodoUsecase {
	TodoRepo := InjectTodoRepository()
	return usecase.NewTodoUsecase(TodoRepo)
}

func InjectTodoFormatter() handler.TodoFormatter {
	return handler.NewTodoFormatter(InjectTodoUsecase())
}
