package injector

import (
	"todo/domain/repository"
	"todo/infra/database"

	"todo/interface/handler"
	"todo/usecase"
)

func InjectDB() database.SqlHandler {
	sqlhandler := database.NewSqlHandler()
	return *sqlhandler
}

func InjectTodoRepository() repository.TodoRepository {
	sqlHandler := InjectDB()
	return database.NewTodoRepository(sqlHandler)
}

func InjectTodoUsecase() usecase.TodoUsecase {
	TodoRepo := InjectTodoRepository()
	return usecase.NewTodoUsecase(TodoRepo)
}

func InjectTodoHandler() handler.TodoHandler {
	return handler.NewTodoHandler(InjectTodoUsecase())
}
