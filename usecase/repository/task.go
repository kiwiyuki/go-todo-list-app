package usecase

import (
	"go-todo-list-app/domain/model"
)

type TaskRepository interface {
	Create(t model.Task) (task model.Task, err error)
	FindAll() (tasks []model.Task, err error)
}
