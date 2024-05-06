package dataModels

import (
	"time"

	"gorm.io/gorm"
)

type Passage struct {
	gorm.Model
	PassageTitle          string
	PassageContent        string
	PassageAuthorUserName string
	PassageAuthorNickName string
	PassageAuthorId       uint
	PassageTag            string
	PassageTime           time.Time `gorm:"default:null"`
}
