package usecase

import (
	"todo/domain/repository"
)

type TodoUsecase struct {
	todoRepo repository.TodoRepository
}
