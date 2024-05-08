package infoModels

type User struct {
	UserId       uint   `json:"userId"`
	UserName     string `json:"userName"`
	UserEmail    string `json:"userEmail"`
	UserNickName string `json:"userNickName"`
	UserIsAdmin  bool   `json:"userIsAdmin"`
}
