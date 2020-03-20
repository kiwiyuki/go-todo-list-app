package database

import (
	"go-todo-list-app/domain/model"
)

type TaskRepository struct {
	Handler SQLHandler
}

func (taskRepository *TaskRepository) Create(t model.Task) (task model.Task, err error) {
	_, err = taskRepository.Handler.Execute(
		"INSERT INTO tasks (title, done) VALUES (?,?)", t.Title, t.Done)
	if err != nil {
		return
	}
	task = t
	return
}

func (taskRepository *TaskRepository) FindAll() (tasks []model.Task, err error) {
	rows, err := taskRepository.Handler.Query("SELECT * FROM tasks")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		task := model.Task{}
		_ = rows.Scan(&task.ID, &task.Title, &task.Done)
		tasks = append(tasks, task)
	}

	return
}
