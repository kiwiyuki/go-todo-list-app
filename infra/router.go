package infra

import (
	"go-todo-list-app/interface/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	taskController := controllers.NewTaskController(NewSQLHandler())

	router.POST("/tasks", func(c *gin.Context) { taskController.Create(c) })
	router.GET("/tasks", func(c *gin.Context) { taskController.Index(c) })
	router.GET("/tasks/:id", func(c *gin.Context) { taskController.Show(c) })
	router.PUT("/tasks/:id", func(c *gin.Context) { taskController.Update(c) })
	Router = router
}
