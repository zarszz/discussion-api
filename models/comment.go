package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ParentInt int64
	Content   string
}
