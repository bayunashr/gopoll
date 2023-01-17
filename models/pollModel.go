package models

import "gorm.io/gorm"

type Poll struct {
	gorm.Model
	Subject     string `gorm:"unique"`
	Description string
	TotalVote   uint
	Visibility  bool
	Archive     bool
	UserID      uint
	PollChoice  []PollChoice
	PollEntry   []PollEntry
}

type PollChoice struct {
	gorm.Model
	Choice    string
	TotalVote uint
	PollID    uint
	PollEntry []PollEntry
}

type PollEntry struct {
	gorm.Model
	UserID       uint
	PollChoiceID uint
	PollID       uint
}
