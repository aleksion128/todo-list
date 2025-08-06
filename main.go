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
	r.GET("/todos/:userId/tasks", handlers.GetTasks)
	r.POST("/todos/:userId/", handlers.CreateTask)
	r.PUT("/todos/:userId/:id", handlers.UpdateTask)
	r.DELETE("/todos/:userId/:id", handlers.DeleteTask)
	r.POST("/users/", handlers.CreateNewUser)
	r.DELETE("/users/delete/:id", handlers.DeleteUserFromBd)
	r.PUT("/users/changeUsername/:id", handlers.UpdateUsernameBd)
	r.PUT("/users/changeEmail/:id", handlers.UpdateEmail)
	r.PUT("/users/changePassword/:id", handlers.UpdatePass)
	r.Run()
}
