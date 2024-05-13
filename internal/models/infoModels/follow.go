package infoModels

type Follow struct {
	ID         uint `json:"id"`
	FromUserId uint `json:"fromUserId"`
	ToUserId   uint `json:"toUserId"`
}
