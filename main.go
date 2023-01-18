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

	r.GET("/", controllers.PublicHome)
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.LogIn)

	r.POST("/poll", middlewares.Auth, controllers.CreatePoll)
	r.GET("/poll", middlewares.Auth, controllers.ReadAllPoll)
	r.GET("/poll/:id", middlewares.Auth, controllers.ReadSpcPoll)
	r.PUT("/poll/:id", middlewares.Auth, controllers.UpdatePoll)
	r.DELETE("/poll/:id", middlewares.Auth, controllers.DeletePoll)

	r.POST("/poll/:id/choice", middlewares.Auth, controllers.CreateChoice)

	r.PUT("/poll/:id/publish", middlewares.Auth, controllers.PublishPoll)
	r.PUT("/poll/:id/archive", middlewares.Auth, controllers.ArchivePoll)
	r.POST("/poll/:id/vote", middlewares.Auth, controllers.VotePoll)

	r.Run()
}
