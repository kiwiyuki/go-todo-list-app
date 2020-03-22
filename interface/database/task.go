package database

import (
	"go-todo-list-app/domain/model"
)

type TaskRepository struct {
	Handler SQLHandler
}

func (taskRepository *TaskRepository) Create(t model.Task) (task model.Task, err error) {
	err = taskRepository.Handler.Create(&t).Error
	task = t
	return
}

func (taskRepository *TaskRepository) FindAll() (tasks []model.Task, err error) {
	err = taskRepository.Handler.Find(&tasks).Error
	return
}

func (taskRepository *TaskRepository) Find(id int64) (task model.Task, err error) {
	err = taskRepository.Handler.First(&task, id).Error
	return
}
