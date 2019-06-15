package infra

import (
	"todo/domain/model"
	"todo/domain/repository"
	"todo/interface/database"
)

type TodoRepository struct {
	database.SqlHandler
}

func NewTodoRepository(sqlHandler database.SqlHandler) repository.TodoRepository {
	todoRepository := TodoRepository{sqlHandler}
	return &todoRepository
}
func (todoRepo *TodoRepository) FindAll() (todos []*model.Todo, err error) {
	rows, err := todoRepo.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		todo := model.Todo{}

		rows.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)

		todos = append(todos, &todo)
	}
	return
}

func (todoRepo *TodoRepository) Find(word string) (todos []*model.Todo, err error) {
	rows, err := todoRepo.Query("SELECT * FROM todos WHERE tasks IN ?", word)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		todo := model.Todo{}

		rows.Scan(&todo.ID, &todo.Task, &todo.LimitDate, &todo.Status)

		todos = append(todos, &todo)
	}
	return
}

func (todoRepo *TodoRepository) Create(todo *model.Todo) (*model.Todo, error) {
	_, err := todoRepo.Execute("INSERT INTO todos (task,limitDate,status) VALUES (?, ?, ?) ", todo.Task, todo.LimitDate, todo.Status)
	return todo, err
}

func (todoRepo *TodoRepository) Update(todo *model.Todo) (*model.Todo, error) {
	_, err := todoRepo.Execute("UPDARE todos SET task = ?,limitDate = ? ,status = ? WHERE id = ?", todo.Task, todo.LimitDate, todo.Status, todo.ID)
	return todo, err
}
