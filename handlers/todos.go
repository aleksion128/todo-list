package handlers

import (
	"log"
	"net/http"
	"strconv"
	"todoList/databas"
	"todoList/models"
	"todoList/storage"

	"github.com/gin-gonic/gin"
)

// Task
func GetTasks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		log.Println(err)
	}
	tasks := storage.GetAll(databas.ConnectDB(), id)
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var tsk models.Task
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		log.Println(err)
		return
	}
	if err := c.BindJSON(&tsk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tsk.UserID = userId
	storage.AddTask(databas.ConnectDB(), tsk)
	c.Status(http.StatusCreated)
}

func UpdateTask(c *gin.Context) {
	var tsk models.Task
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		log.Println(err)
		return
	}
	if err := c.BindJSON(&tsk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tsk.UserID = userId
	storage.UpdateTask(databas.ConnectDB(), tsk)
	c.Status(http.StatusOK)
}

func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	storage.DeleteTask(databas.ConnectDB(), id)
	c.Status(http.StatusOK)
}

// User
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
	var user models.User
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
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	storage.UpdateEmail(databas.ConnectDB(), id, user.Email)
	c.Status(200)
}

func UpdatePass(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	storage.UpdatePassword(databas.ConnectDB(), id, user.Password)
	c.Status(200)
}
