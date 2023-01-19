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

func UpdatePoll(c *gin.Context) {
	id := c.Param("id")
	curUser, _ := c.Get("currentUser")
	var curPoll models.Poll
	initializers.DB.Where("id", id).Take(&curPoll)
	if curPoll.UserID == curUser.(models.User).ID {
		if !curPoll.Visibility && !curPoll.Archive {
			var poll struct {
				Subject     string
				Description string
			}
			if c.Bind(&poll) != nil {
				c.JSON(400, gin.H{
					"message": "error, fail to read body",
				})
			} else {
				result := initializers.DB.Model(&curPoll).Updates(models.Poll{Subject: poll.Subject, Description: poll.Description})
				if result.Error != nil {
					c.JSON(400, gin.H{
						"message": "error, fail to read body",
					})
				} else {
					c.JSON(200, gin.H{
						"message": "success, poll update",
					})
				}
			}
		} else {
			c.JSON(400, gin.H{
				"message": "error, no poll",
			})
		}
	} else {
		c.JSON(400, gin.H{
			"message": "error, no poll",
		})
	}
}

func DeletePoll(c *gin.Context) {
	id := c.Param("id")
	curUser, _ := c.Get("currentUser")
	var curPoll models.Poll
	initializers.DB.Where("id", id).Take(&curPoll)
	if curPoll.UserID == curUser.(models.User).ID {
		deletedPoll := initializers.DB.Delete(&curPoll)
		var curPollChoice []models.PollChoice
		deletedChoice := initializers.DB.Where("poll_id = ?", curPoll.ID).Delete(&curPollChoice)
		if deletedPoll.Error != nil || deletedChoice.Error != nil {
			c.JSON(400, gin.H{
				"message": "error, fail to delete poll",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "success, poll gone",
			})
		}
	} else {
		c.JSON(400, gin.H{
			"message": "error, no poll",
		})
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
	curUser, _ := c.Get("currentUser")
	var spcPoll models.Poll
	pollResult := initializers.DB.Where("id", id).Preload("PollChoice").Take(&spcPoll)
	if spcPoll.UserID == curUser.(models.User).ID {
		if pollResult.Error != nil {
			c.JSON(400, gin.H{
				"message": "error, poll didnt exist",
			})
		} else {
			c.JSON(200, gin.H{
				"poll": spcPoll,
			})
		}
	} else {
		if !spcPoll.Visibility {
			c.JSON(400, gin.H{
				"message": "error, poll is private",
			})
		} else {
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
	}
}

func CreateChoice(c *gin.Context) {
	id := c.Param("id")
	curUser, _ := c.Get("currentUser")
	var curPoll models.Poll
	initializers.DB.Take(&curPoll, id)
	if curPoll.UserID == curUser.(models.User).ID {
		if !curPoll.Visibility && !curPoll.Archive {
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
		} else {
			c.JSON(400, gin.H{
				"message": "error, poll is either on air or archived",
			})
		}
	} else {
		c.JSON(400, gin.H{
			"message": "error, poll is not yours",
		})
	}
}

func DeleteChoice(c *gin.Context) {
	id, ch := c.Param("id"), c.Param("ch")
	var curChoice models.PollChoice
	result := initializers.DB.Where("id = ? AND poll_id = ?", ch, id).Delete(&curChoice)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "error, fail to delete choice",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success, choice deleted",
		})
	}
}

func PublishPoll(c *gin.Context) {
	id := c.Param("id")
	curUser, _ := c.Get("currentUser")
	var curPoll models.Poll
	initializers.DB.Take(&curPoll, id)
	if uint(curUser.(models.User).ID) == curPoll.UserID {
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
	} else {
		c.JSON(401, gin.H{
			"message": "error, youre not the owner",
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
		c.JSON(401, gin.H{
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
