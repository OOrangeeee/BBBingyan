package dataModels

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	FromUserId uint
	ToUserId   uint
}
