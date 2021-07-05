package controllers

import (
	"net/http"
	"todo-app/utils/context"

	"github.com/gin-gonic/gin"
)

func Login(c *context.Context) {
	c.HTML(http.StatusOK, "pages/login", gin.H{})
}

func GoogleOAuth(c *context.Context) {
	// TODO update session
	c.Redirect(http.StatusFound, "/")
}
