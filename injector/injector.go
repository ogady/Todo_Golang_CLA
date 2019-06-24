package injector

import (
	"todo/domain/repository"
	"todo/handler"
	"todo/infra"
	"todo/usecase"
)

func InjectDB() infra.SqlHandler {
	sqlhandler := infra.NewSqlHandler()
	return *sqlhandler
}

func InjectTodoRepository() repository.TodoRepository {
	sqlHandler := InjectDB()
	return infra.NewTodoRepository(sqlHandler)
}

func InjectTodoUsecase() usecase.TodoUsecase {
	TodoRepo := InjectTodoRepository()
	return usecase.NewTodoUsecase(TodoRepo)
}

func InjectTodoHandler() handler.TodoHandler {
	return handler.NewTodoHandler(InjectTodoUsecase())
}
