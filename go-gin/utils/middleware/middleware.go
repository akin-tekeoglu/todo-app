package middleware

import (
	"net/http"
	"todo-app/models/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("userId") == nil {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
	userId := session.Get("userId").(int)
	c.Set("user", user.GetById(userId))
	c.Next()
}
