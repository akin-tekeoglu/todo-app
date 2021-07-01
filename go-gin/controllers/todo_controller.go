package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupTodoControllers(e *gin.Engine) {

	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{})
	})
}
