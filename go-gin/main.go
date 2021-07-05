package main

import (
	"todo-app/controllers"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.HTMLRender = ginview.Default()
	r.GET("/", controllers.ShowAll)
	r.GET("/todo/form/*id", controllers.ShowOne)
	r.POST("/todo/form/*id", controllers.Save)
	r.POST("/todo/:id/toggle", controllers.ToggleCompleted)
	r.Run()
}
