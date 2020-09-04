package models

import (
	"gorm.io/gorm"
)

type Discussion struct {
	gorm.Model
	Title   string
	Content string
	User    User
}
