package main

import (
	"go-todo-list-app/infra"
)

func main() {
	err := infra.Router.Run()
	if err != nil {
		return
	}
}
