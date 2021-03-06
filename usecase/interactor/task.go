package usecase

import (
	"go-todo-list-app/domain/model"
	repository "go-todo-list-app/usecase/repository"
)

type TaskInteractor struct {
	TaskRepository repository.TaskRepository
}

func (tasksInteractor *TaskInteractor) Create(t model.Task) (task model.Task, err error) {
	created, err := tasksInteractor.TaskRepository.Create(t)
	if err != nil {
		return
	}
	task = created
	return
}

func (tasksInteractor *TaskInteractor) FindAll() (tasks []model.Task, err error) {
	tasks, err = tasksInteractor.TaskRepository.FindAll()
	return
}

func (tasksInteractor *TaskInteractor) Find(id int64) (task model.Task, err error) {
	task, err = tasksInteractor.TaskRepository.Find(id)
	return
}

func (TaskInteractor *TaskInteractor) Update(t model.Task) (err error) {
	err = TaskInteractor.TaskRepository.Update(t)
	return
}
