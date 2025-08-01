package main

import (
	"todoList/handlers"

	"github.com/gin-gonic/gin"
)

var ID int

func main() {
	r := gin.Default()
	r.GET("/todos", handlers.GetTasks)
	r.POST("/todos", handlers.CreateTask)
	r.PUT("/todos/:id", handlers.UpdateTask)
	r.DELETE("/todos/:id", handlers.DeleteTask)

	r.Run()
}
