package main

import "BBBingyan/internal/utils"

func main() {
	confirmTokenTool := utils.ConfirmTokenTool{}
	print(confirmTokenTool.GenerateConfirmToken())

}
