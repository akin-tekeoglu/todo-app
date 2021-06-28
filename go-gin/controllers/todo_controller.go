package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupTodoControllers(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"message": "Todo app",
		})
	})
}
