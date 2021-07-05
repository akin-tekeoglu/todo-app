package middleware

import (
	"net/http"
	"todo-app/models/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	if c.FullPath() == "/login" || c.FullPath() == "/oauth/google" {
		return
	}
	session := sessions.Default(c)
	if session.Get("userId") == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	userId := session.Get("userId").(int)
	c.Set("user", user.GetById(userId))
}

func CurrentUser(c *gin.Context) *user.User {
	currentUser, exists := c.Get("user")
	if !exists {
		return nil
	}
	return currentUser.(*user.User)
}
