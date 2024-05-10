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
	PassageBeLikedCount   int
	PassageCommentCount   int
	PassageTime           time.Time `gorm:"default:null"`
}
