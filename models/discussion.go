package models

import (
	"gorm.io/gorm"
)

type Discussion struct {
	gorm.Model
	Title      string
	Content    string
	CategoryID uint
	UserID     uint
	User       User
	Comments   []Comment `gorm:"foreignKey:DiscussionID"`
}
