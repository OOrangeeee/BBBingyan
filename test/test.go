package main

import (
	"BBBingyan/internal/utils"
	"fmt"
)

func main() {
	matchTool := utils.MatchTool{}
	atUsers := matchTool.MatchAtInString("Hello, @Alice,hello,w ,,@Bob ")
	for _, user := range atUsers {
		fmt.Println(user)
	}

}
