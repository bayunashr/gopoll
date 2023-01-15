package models

import "gorm.io/gorm"

type Poll struct {
	gorm.Model
	Subject     string `gorm:"unique;not null"`
	Description string
	TotalVote   int
	UserID      int
	User        User
}

type PollChoice struct {
	gorm.Model
	Choice    string `gorm:"not null"`
	TotalVote int
	PollID    int
	Poll      Poll
}

type PollEntry struct {
	gorm.Model
	UserID       int
	PollChoiceID int
	PollID       int
	User         User
	PollChoice   PollChoice
	Poll         Poll
}
