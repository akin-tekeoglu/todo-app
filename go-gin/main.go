package main

import (
	"todo-app/controllers"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.HTMLRender = ginview.Default()
	r.GET("/", controllers.ShowAll)
	r.POST("/todo/form", controllers.Create)
	r.Run()
}
