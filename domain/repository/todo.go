package repository

import (
	"todo/domain/model"

)

//TodoRepository is interface for infrastructure
type TodoRepository interface {
	FindAll() (todos []*model.Todo, err error)
	Find(word string) (todos []*model.Todo, err error)
	Create(todo *model.Todo) (*model.Todo, error)
	Update(todo *model.Todo) (*model.Todo, error)
}


