package models

import "gorm.io/gorm"

type Poll struct {
	gorm.Model
	Subject     string `gorm:"unique;not null"`
	Description string
	TotalVote   int
	OwnedBy     int
	UserOwner   User `gorm:"foreignKey:OwnedBy"`
}

type PollChoice struct {
	gorm.Model
	Choice    string `gorm:"not null"`
	TotalVote int
	OwnedBy   int
	PollOwner Poll `gorm:"foreignKey:OwnedBy"`
}

type PollEntry struct {
	gorm.Model
	VoteBy      string
	VoteOn      int
	VoteOwnedBy int
	UserVoted   User       `gorm:"foreignKey:VoteBy"`
	ChoiceVoted PollChoice `gorm:"foreignKey:VoteOn"`
	PollOwner   Poll       `gorm:"foreignKey:VoteOwnedBy"`
}
