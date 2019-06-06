package handler

type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}
