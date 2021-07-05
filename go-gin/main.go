package main

import (
	"os"
	"todo-app/controllers"
	"todo-app/utils/context"
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
	e := context.Engine{GinEngine: gin.Default()}
	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET")))
	e.Use(sessions.Sessions("user", store))
	e.GinEngine.HTMLRender = ginview.Default()
	e.GET("/login", controllers.Login)
	e.GET("/oauth/google", controllers.GoogleOAuth)
	e.Use(middleware.Authenticate)
	{
		e.GET("/", controllers.ShowAll)
		e.GET("/todo/form/*id", controllers.ShowOne)
		e.POST("/todo/form/*id", controllers.Save)
		e.POST("/todo/:id/toggle", controllers.ToggleCompleted)
	}
	e.GinEngine.Run()
}
