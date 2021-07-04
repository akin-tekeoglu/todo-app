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
	r.GET("/todo/form/*id", controllers.ShowOne)
	r.POST("/todo/form/*id", controllers.Save)
	r.POST("/todo/:id/toggle", controllers.ToggleCompleted)
	r.Run()
}
