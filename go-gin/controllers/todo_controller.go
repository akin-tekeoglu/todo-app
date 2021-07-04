package controllers

import (
	"net/http"
	"strconv"
	"todo-app/models/todo"

	"github.com/gin-gonic/gin"
)

func ShowAll(c *gin.Context) {
	todos := todo.GetAll()
	c.HTML(http.StatusOK, "pages/index", gin.H{"todos": todos})
}

func ShowOne(c *gin.Context) {
	todoItem := new(todo.Todo)
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		*todoItem = todo.GetById(id)
	}
	c.HTML(http.StatusOK, "todo", gin.H{"todo": todoItem})
}

func Save(c *gin.Context) {
	todoItem, result := todo.New(c.PostForm("title"), c.PostForm("description"), c.PostForm("eventDate"))
	if result != todo.Ok {
		c.HTML(http.StatusOK, "todo", gin.H{"result": result, "title": c.PostForm("title"), "description": c.PostForm("description"), "eventDate": c.PostForm("eventDate")})
	}
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		todoItem.Id = id
	}
	todoItem.Save()
	c.HTML(http.StatusOK, "todo_success", gin.H{})
}

func ToggleCompleted(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		todoItem := todo.GetById(id)
		todoItem.ToggleCompleted()
		c.HTML(http.StatusOK, "todo_success", gin.H{})
	} else {
		c.Status(http.StatusNotFound)
	}
}
