package controllers

import (
	"github.com/bayunashr/gopoll/initializers"
	"github.com/bayunashr/gopoll/models"
	"github.com/gin-gonic/gin"
)

func CreatePoll(c *gin.Context) {
	var poll struct {
		Subject     string
		Description string
	}
	if c.Bind(&poll) != nil {
		c.JSON(400, gin.H{
			"message": "error, fail to read body",
		})
	} else {
		if poll.Subject == "" {
			c.JSON(400, gin.H{
				"message": "error, subject is mandatory",
			})
		} else {
			curUser, _ := c.Get("currentUser")
			newPoll := models.Poll{Subject: poll.Subject, Description: poll.Description, TotalVote: 0, UserID: int(curUser.(models.User).ID)}
			result := initializers.DB.Create(&newPoll)
			if result.Error != nil {
				c.JSON(400, gin.H{
					"message": "error, fail to create new poll",
				})
			} else {
				c.JSON(200, gin.H{
					"message": "success, created new poll",
				})
			}
		}
	}
}
