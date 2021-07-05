package controllers

import (
	"net/http"
	"strconv"
	"todo-app/models/todo"
	"todo-app/utils/context"

	"github.com/gin-gonic/gin"
)

func ShowAll(c *context.Context) {
	todos := todo.GetAll(c.User.Id)
	c.HTML(http.StatusOK, "pages/index", gin.H{"todos": todos})
}

func ShowOne(c *context.Context) {
	todoItem := todo.Todo{}
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		todoItem = todo.GetById(id, c.User.Id)
	}
	c.HTML(http.StatusOK, "pages/todo", gin.H{"todo": todoItem})
}

func Save(c *context.Context) {
	todoItem, result := todo.New(c.PostForm("title"), c.PostForm("description"), c.PostForm("eventDate"))
	if result != todo.Ok {
		c.HTML(http.StatusOK, "pages/todo", gin.H{"result": result, "title": c.PostForm("title"), "description": c.PostForm("description"), "eventDate": c.PostForm("eventDate")})
	}
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		todoItem.Id = id
	}
	todoItem.UserId = c.User.Id
	todoItem.Save()
	c.HTML(http.StatusOK, "pages/todoSuccess", gin.H{})
}

func ToggleCompleted(c *context.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		todoItem := todo.GetById(id, c.User.Id)
		todoItem.ToggleCompleted()
		c.HTML(http.StatusOK, "pages/todoSuccess", gin.H{})
	} else {
		c.Status(http.StatusNotFound)
	}
}
