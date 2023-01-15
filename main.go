package main

import (
	"github.com/bayunashr/gopoll/controllers"
	"github.com/bayunashr/gopoll/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.LoadDb()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "you are home",
		})
	})
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.LogIn)
	r.Run()
}
