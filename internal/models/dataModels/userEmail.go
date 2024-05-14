package dataModels

import (
	"time"

	"gorm.io/gorm"
)

type UserEmail struct {
	gorm.Model
	Email                   string    `gorm:"unique"`
	EmailLastSentOfRegister time.Time `gorm:"default:null"`
	EmailLastSentOfLogin    time.Time `gorm:"default:null"`
}
