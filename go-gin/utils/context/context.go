package context

import (
	"todo-app/models/user"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type Context struct {
	ginContext *gin.Context
	User       *user.User
}

func newContext(c *gin.Context) Context {
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
	ginEngine *gin.Engine
}

func NewEngine(e *gin.Engine) Engine {
	return Engine{ginEngine: e}
}

func (e *Engine) GET(relativePath string, handler func(c *Context)) {
	e.ginEngine.GET(relativePath, func(c *gin.Context) {
		ctx := newContext(c)
		handler(&ctx)
	})
}

func (e *Engine) POST(relativePath string, handler func(c *Context)) {
	e.ginEngine.POST(relativePath, func(c *gin.Context) {
		ctx := newContext(c)
		handler(&ctx)
	})
}

func (e *Engine) Use(middleware gin.HandlerFunc) {
	e.ginEngine.Use(middleware)
}

func (e *Engine) Run(addr ...string) {
	e.ginEngine.Run(addr...)
}

func (e *Engine) HTMLRender(htmlRender render.HTMLRender) {
	e.ginEngine.HTMLRender = htmlRender
}
