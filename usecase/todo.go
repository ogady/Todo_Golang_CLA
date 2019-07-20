package usecase

import (
	"todo/domain/model"
	"todo/domain/repository"
)

/*
TodoUsecase DIのためのtodoUsecase の interface
*/
type TodoUsecase interface {
	Search(string) (todo []*model.Todo, err error)
	View() (todo []*model.Todo, err error)
	Add(*model.Todo) (err error)
	Edit(*model.Todo) (err error)
}

/*
todoUsecase Todoに関するユースケース
interface層に依存される
domein層に依存する
*/
type todoUsecase struct {
	todoRepo repository.TodoRepository
}

/*
NewTodoUsecase is constructor for init
TodoUsecase interface を返却している
go言語において呼び出し元にinterfaceのみを公開することで実装の詳細を隠ぺいすると、gomockを使ったテストがやり易くなる。
NewTodoUsecase 関数で todoUsecase(非公開) のメモリを確保し、interfaceである TodoUsecase（公開） 型として利用者に公開します。
*/
func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	todoUsecase := todoUsecase{todoRepo: todoRepo}
	return &todoUsecase
}

// Search 入力された内容でTodoを検索する
func (usecase *todoUsecase) Search(word string) (todo []*model.Todo, err error) {
	todo, err = usecase.todoRepo.Find(word)
	return
}

// View はTodoの一覧を表示する
func (usecase *todoUsecase) View() (todo []*model.Todo, err error) {
	todo, err = usecase.todoRepo.FindAll()
	return
}

// Add はTodoを新規追加する
func (usecase *todoUsecase) Add(todo *model.Todo) (err error) {
	_, err = usecase.todoRepo.Create(todo)
	return
}

// Edit はTodoを編集する
func (usecase *todoUsecase) Edit(todo *model.Todo) (err error) {
	_, err = usecase.todoRepo.Update(todo)
	return
}
