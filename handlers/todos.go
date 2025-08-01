package handlers

import (
	"net/http"
	"strconv"
	"todoList/models"
	"todoList/storage"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks := storage.GetAll()
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var tsk models.Task
	if err := c.BindJSON(&tsk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	storage.AddTask(tsk)
	c.Status(http.StatusCreated)
}

func UpdateTask(c *gin.Context) {
	var tsk models.Task
	if err := c.BindJSON(&tsk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	storage.ChangeTask(tsk)
	c.Status(http.StatusOK)
}

func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	storage.DeleteTask(id)
	c.Status(http.StatusOK)
}
