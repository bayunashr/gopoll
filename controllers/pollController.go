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
				"message": "error, fill all required fields",
			})
		} else {
			curUser, _ := c.Get("currentUser")
			newPoll := models.Poll{Subject: poll.Subject, Description: poll.Description, TotalVote: 0, UserID: uint(curUser.(models.User).ID)}
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

func ReadMyPoll(c *gin.Context) {
	curUser, _ := c.Get("currentUser")
	var myPoll []models.Poll
	result := initializers.DB.Where("user_id = ?", curUser.(models.User).ID).Find(&myPoll)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "error, poll didnt exist",
		})
	} else {
		c.JSON(200, gin.H{
			"poll": myPoll,
		})
	}
}

func ReadAllPoll(c *gin.Context) {
	var allPoll []models.Poll
	result := initializers.DB.Find(&allPoll)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "error, poll didnt exist",
		})
	} else {
		c.JSON(200, gin.H{
			"poll": allPoll,
		})
	}
}

func ReadSpcPoll(c *gin.Context) {
	id := c.Param("id")
	var spcPoll []models.Poll
	var spcPollChoice []models.PollChoice
	pollResult := initializers.DB.Where("id", id).First(&spcPoll)
	pollChoiceResult := initializers.DB.Where("poll_id", id).Find(&spcPollChoice)
	if pollResult.Error != nil || pollChoiceResult.Error != nil {
		c.JSON(400, gin.H{
			"message": "error, poll didnt exist",
		})
	} else {
		c.JSON(200, gin.H{
			"poll":       spcPoll,
			"pollChoice": spcPollChoice,
		})
	}
}

func CreateChoice(c *gin.Context) {
	var choice struct {
		Choice string
	}
	if c.Bind(&choice) != nil {
		c.JSON(400, gin.H{
			"message": "error, fail to read body",
		})
	} else {
		if choice.Choice == "" {
			c.JSON(400, gin.H{
				"message": "error, fill all required fields",
			})
		} else {
			id := c.Param("id")
			var curPoll models.Poll
			initializers.DB.First(&curPoll, id)
			newChoice := models.PollChoice{Choice: choice.Choice, TotalVote: 0, PollID: uint(curPoll.ID)}
			result := initializers.DB.Create(&newChoice)
			if result.Error != nil {
				c.JSON(400, gin.H{
					"message": "error, fail to create new poll choice",
				})
			} else {
				c.JSON(200, gin.H{
					"message": "success, created new poll choice",
				})
			}
		}
	}
}

func ReadChoice(c *gin.Context) {
	id := c.Param("id")
	var myChoice []models.PollChoice
	result := initializers.DB.Where("poll_id = ?", id).Find(&myChoice)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "error, choice didnt exist",
		})
	} else {
		c.JSON(200, gin.H{
			"choice": myChoice,
		})
	}
}
