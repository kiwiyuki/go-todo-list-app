package controllers

import (
	"go-todo-list-app/domain/model"
	"go-todo-list-app/interface/database"
	usecase "go-todo-list-app/usecase/interactor"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (taskController *TaskController) Show(c Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	task, err := taskController.Interactor.Find(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, task)
}

func (taskController *TaskController) Update(c Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	_, err := taskController.Interactor.Find(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}

	t := model.Task{}
	err = c.Bind(&t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}

	t.ID = uint(id)
	err = taskController.Interactor.Update(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
