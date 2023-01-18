package controllers

import (
	"github.com/bayunashr/gopoll/initializers"
	"github.com/bayunashr/gopoll/models"
	"github.com/gin-gonic/gin"
)

func PublicHome(c *gin.Context) {
	var hotPoll []models.Poll
	result := initializers.DB.Where("visibility = true AND archive = false").Order("total_vote desc").Limit(10).Find(&hotPoll)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "error, fail to read poll",
		})
	} else {
		c.JSON(200, gin.H{
			"data": hotPoll,
		})
	}
}
