package main

import (
	"github.com/bayunashr/gopoll/controllers"
	"github.com/bayunashr/gopoll/initializers"
	"github.com/bayunashr/gopoll/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.LoadDb()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()
	r.GET("/", controllers.GuestHome)
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.LogIn)
	r.POST("/create", middlewares.Auth, controllers.CreatePoll)
	r.POST("/create/:id/choice", middlewares.Auth, controllers.CreateChoice)
	r.Run()
}
