package dataModels

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	FromUserId  uint
	ToPassageId uint
}
