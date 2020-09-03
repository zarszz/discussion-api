package models

import (
	"gorm.io/gorm"
)

type Discussions struct {
	gorm.Model
	Title   string
	Content string
}
