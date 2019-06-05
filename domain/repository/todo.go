package repository

import (
	"todo/domain/model"
)

//TodoRepository is interface for infrastructure
type TodoRepository interface {
	Find(string) ([]*model.Todo, error)
	FindAll() ([]*model.Todo, error)
	Create(*model.Todo) (*model.Todo, error)
	Update(*model.Todo) (*model.Todo, error)
}
