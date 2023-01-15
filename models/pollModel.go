package models

import "gorm.io/gorm"

type Poll struct {
	gorm.Model
	Subject     string `gorm:"unique;not null"`
	Description string
	TotalVote   uint
	UserID      uint
}

type PollChoice struct {
	gorm.Model
	Choice    string `gorm:"not null"`
	TotalVote uint
	PollID    uint
}

type PollEntry struct {
	gorm.Model
	UserID       uint
	PollChoiceID uint
	PollID       uint
}
