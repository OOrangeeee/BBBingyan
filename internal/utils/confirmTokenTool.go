package utils

import (
	"math/rand"
	"time"
)

type ConfirmTokenTool struct {
}

func (cT *ConfirmTokenTool) GenerateConfirmToken() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	confirmToken := make([]rune, 6)
	for i := 0; i < 6; i++ {
		confirmToken[i] = rune(r.Intn(10) + 48)
	}
	return string(confirmToken)
}
