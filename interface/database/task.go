package database

import (
	"go-todo-list-app/domain/model"
)

type TaskRepository struct {
	Handler SQLHandler
}

func (taskRepository *TaskRepository) Create(t model.Task) (task model.Task, err error) {
	result, err := taskRepository.Handler.Execute(
		"INSERT INTO tasks (title, done) VALUES (?,?)", t.Title, t.Done)
	if err != nil {
		return
	}
	task = t
	task.ID, _ = result.LastInsertId()
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
		err = rows.Scan(&task.ID, &task.Title, &task.Done)
		if err != nil {
			return
		}
		tasks = append(tasks, task)
	}

	return
}
