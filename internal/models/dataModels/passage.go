package dataModels

import (
	"time"

	"gorm.io/gorm"
)

type passage struct {
	gorm.Model
	PassageTitle   string
	PassageContent string
	PassageAuthor  string
	PassageTags    string
	PassageTime    time.Time `gorm:"default:null"`
}
