package handlers

import (
	"net/http"
	"strconv"
	"todoList/databas"
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

func CreateNewUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	storage.CreateUser(databas.ConnectDB(), user.Username, user.Email, user.Password)
	c.Status(http.StatusOK)
}

func DeleteUserFromBd(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	storage.DeleteUser(databas.ConnectDB(), id)
	c.Status(http.StatusOK)
}

func UpdateUsernameBd(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.UpdatedUsernameUser
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	storage.UpdateUsername(databas.ConnectDB(), id, user.Username)
	c.Status(http.StatusOK)
}

func UpdateEmail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.UpdatedEmailUser
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	storage.UpdateEmail(databas.ConnectDB(), id, user.Email)
}
