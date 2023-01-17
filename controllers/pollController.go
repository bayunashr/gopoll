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
			newPoll := models.Poll{Subject: poll.Subject, Description: poll.Description, TotalVote: 0, Visibility: false, Archive: false, UserID: uint(curUser.(models.User).ID)}
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

func ReadAllPoll(c *gin.Context) {
	isMine := c.Query("mine")
	var allPoll []models.Poll
	if isMine == "yes" {
		curUser, _ := c.Get("currentUser")
		result := initializers.DB.Where("user_id", curUser.(models.User).ID).Preload("PollChoice").Find(&allPoll)
		if result.Error != nil {
			c.JSON(400, gin.H{
				"message": "error, poll didnt exist",
			})
		} else {
			c.JSON(200, gin.H{
				"poll": allPoll,
			})
		}
	} else {
		result := initializers.DB.Where("visibility = true AND archive = false").Preload("PollChoice").Find(&allPoll)
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
}

func ReadSpcPoll(c *gin.Context) {
	id := c.Param("id")
	var spcPoll []models.Poll
	pollResult := initializers.DB.Where("id", id).Preload("PollChoice").First(&spcPoll)
	if pollResult.Error != nil {
		c.JSON(400, gin.H{
			"message": "error, poll didnt exist",
		})
	} else {
		c.JSON(200, gin.H{
			"poll": spcPoll,
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

func PublishPoll(c *gin.Context) {
	id := c.Param("id")
	var curPoll models.Poll
	initializers.DB.Take(&curPoll, id)
	result := initializers.DB.Model(&curPoll).Update("visibility", true)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "error, fail to publish poll",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success, poll on air",
		})
	}
}

func ArchivePoll(c *gin.Context) {
	id := c.Param("id")
	curUser, _ := c.Get("currentUser")
	var curPoll models.Poll
	initializers.DB.Take(&curPoll, id)
	if uint(curUser.(models.User).ID) == curPoll.UserID {
		result := initializers.DB.Model(&curPoll).Update("archive", true)
		if result.Error != nil {
			c.JSON(400, gin.H{
				"message": "error, fail to archive poll",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "success, poll archived",
			})
		}
	} else {
		c.JSON(400, gin.H{
			"message": "error, youre not the owner",
		})
	}
}

func VotePoll(c *gin.Context) {
	id := c.Param("id")
	curUser, _ := c.Get("currentUser")
	var curPoll models.Poll
	initializers.DB.First(&curPoll, id)
	var pollEntry struct {
		PollChoiceID uint
	}
	c.Bind(&pollEntry)
	newVote := models.PollEntry{UserID: uint(curUser.(models.User).ID), PollChoiceID: pollEntry.PollChoiceID, PollID: curPoll.ID}
	result := initializers.DB.Create(&newVote)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "error, fail to vote",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success, vote hit",
		})
	}
}
