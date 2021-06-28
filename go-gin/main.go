package main

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	controllers.SetupTodoControllers(r)
	r.Run()
}
