package controllers

import (
	"net/http"
	"strconv"
	"time"
	"todo-app/models/todo"

	"github.com/gin-gonic/gin"
)

func ShowAll(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{})
}

func ShowOne(c *gin.Context) {
	todoItem := new(todo.Todo)
	if id, err := strconv.Atoi(c.Param("id")); err == nil {
		todoItem = todo.GetById(id)
	}
	c.HTML(http.StatusOK, "todo", gin.H{"todo": todoItem})
}

func Create(c *gin.Context) {
	eventDate, _ := time.Parse(time.RFC3339, c.PostForm("eventDate"))
	todo := todo.Todo{
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		EventDate:   eventDate,
	}
	todo.Save()
	c.HTML(http.StatusOK, "index", gin.H{})
}
