package infoModels

type Follow struct {
	ID       uint `json:"id"`
	FromUser User `json:"fromUser"`
	ToUser   User `json:"toUser"`
}
