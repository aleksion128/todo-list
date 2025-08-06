package main

import (
	"todoList/databas"
	"todoList/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var ID int

func main() {
	db := databas.ConnectDB()
	defer db.Close()

	r := gin.Default()
	r.GET("/todos", handlers.GetTasks)
	r.POST("/todos", handlers.CreateTask)
	r.PUT("/todos/:id", handlers.UpdateTask)
	r.DELETE("/todos/:id", handlers.DeleteTask)
	r.POST("/users", handlers.CreateNewUser)
	r.DELETE("/users/delete/:id", handlers.DeleteUserFromBd)
	r.PUT("/users/changeUsername/:id", handlers.UpdateUsernameBd)
	r.Run()
}
