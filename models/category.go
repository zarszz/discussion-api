package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Category string
}
