package controllers

import (
	"go-todo-list-app/domain/model"
	"go-todo-list-app/interface/database"
	usecase "go-todo-list-app/usecase/interactor"
	"net/http"
)

type TaskController struct {
	Interactor usecase.TaskInteractor
}

func NewTaskController(handler database.SQLHandler) *TaskController {
	return &TaskController{
		Interactor: usecase.TaskInteractor{
			TaskRepository: &database.TaskRepository{
				Handler: handler,
			},
		},
	}
}

func (taskController *TaskController) Create(c Context) {
	t := model.Task{}
	err := c.Bind(&t)

	task, err := taskController.Interactor.Create(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (taskController *TaskController) Index(c Context) {
	tasks, err := taskController.Interactor.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, tasks)
}
