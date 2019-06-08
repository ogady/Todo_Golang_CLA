package infra

import (
	"todo/domain/model"
	"todo/interface/database"
)

type TodoRepository struct {
	database.SqlHandler
}

func (todoRepo *TodoRepository) FindAll() (todos []*model.Todo, err error) {
	rows, err := todoRepo.Query("")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var task string
		var limitDate string
		var status bool
		if err := rows.Scan(&id, &task, &limitDate, &status); err != nil {
			continue
		}
		todo := model.Todo{}
		todos = append(todos, &todo)
	}
	return
}

func (todoRepo *TodoRepository) Find(string) (todos []*model.Todo, err error) {
	return
}

func (todoRepo *TodoRepository) Create(*model.Todo) (todo *model.Todo, err error) {
	return
}

func (todoRepo *TodoRepository) Update(*model.Todo) (todo *model.Todo, err error) {
	return
}
