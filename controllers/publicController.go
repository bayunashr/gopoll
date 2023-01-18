package controllers

import "github.com/gin-gonic/gin"

func PublicHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success, this is home",
	})
}
