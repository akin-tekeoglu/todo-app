package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/login", gin.H{})
}

func GoogleOAuth(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
