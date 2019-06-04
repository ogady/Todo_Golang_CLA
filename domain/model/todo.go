package model

import "time"

//Todo is TodoModel
type Todo struct {
	id        int
	task      string
	limitDate string
	status    bool
}

//NewTodo is constructor
func NewTodo(id int, task string) *Todo {
	todo := Todo{id: id, task: task, limitDate: time.Now().Format("0000-00-00"), status: false}
	return &todo
}
