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

	r.POST("/poll", middlewares.Auth, controllers.CreatePoll)
	r.GET("/poll", middlewares.Auth, controllers.ReadAllPoll)
	r.GET("/poll/:id", middlewares.Auth, controllers.ReadSpcPoll)
	r.GET("/poll/mine", middlewares.Auth, controllers.ReadMyPoll)
	r.POST("/poll/:id", middlewares.Auth, controllers.CreateChoice)
	r.Run()
}
