package dataModels

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CommentContent string
	FromUserId     uint
	ToPassageId    uint
}
