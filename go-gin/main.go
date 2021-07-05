package main

import (
	"os"
	"todo-app/controllers"
	"todo-app/utils/middleware"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET")))
	r.Use(sessions.Sessions("user", store))
	r.Use(middleware.Authenticate)
	r.HTMLRender = ginview.Default()
	r.GET("/", controllers.ShowAll)
	r.GET("/login", controllers.Login)
	r.GET("/oauth/google", controllers.GoogleOAuth)
	r.GET("/todo/form/*id", controllers.ShowOne)
	r.POST("/todo/form/*id", controllers.Save)
	r.POST("/todo/:id/toggle", controllers.ToggleCompleted)
	r.Run()
}
