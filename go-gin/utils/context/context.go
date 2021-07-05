package context

import (
	"todo-app/models/user"

	"github.com/gin-gonic/gin"
)

type Context struct {
	ginContext *gin.Context
	User       *user.User
}

func New(c *gin.Context) Context {
	wrappers := Context{ginContext: c}
	currentUser, exists := c.Get("user")
	if exists {
		wrappers.User = currentUser.(*user.User)
	}
	return wrappers
}

func (c *Context) Redirect(code int, location string) {
	c.ginContext.Redirect(code, location)
}

func (c *Context) Status(code int) {
	c.ginContext.Status(code)
}

func (c *Context) HTML(code int, name string, h gin.H) {
	if c.User != nil {
		h["user"] = c.User
	}
	c.ginContext.HTML(code, name, h)
}

func (c *Context) Param(key string) string {
	return c.ginContext.Param(key)
}

func (c *Context) PostForm(key string) string {
	return c.ginContext.PostForm(key)
}

type Engine struct {
	GinEngine *gin.Engine
}

func (e *Engine) GET(relativePath string, handler func(c *Context)) {
	e.GinEngine.GET(relativePath, func(c *gin.Context) {
		ctx := New(c)
		handler(&ctx)
	})
}

func (e *Engine) POST(relativePath string, handler func(c *Context)) {
	e.GinEngine.POST(relativePath, func(c *gin.Context) {
		ctx := New(c)
		handler(&ctx)
	})
}

func (e *Engine) Use(middleware gin.HandlerFunc) {
	e.GinEngine.Use(middleware)
}
