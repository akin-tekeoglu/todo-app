package main

import (
	"todo-app/controllers"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.HTMLRender = ginview.Default()
	controllers.SetupTodoControllers(r)
	r.Run()
}
