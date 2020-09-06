package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ParentID     uint `gorm:"index"`
	Content      string
	DiscussionID uint
	UserID       uint
}
