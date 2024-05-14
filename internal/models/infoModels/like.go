package infoModels

type Like struct {
	ID        uint         `json:"id"`
	FromUser  User         `json:"fromUser"`
	ToPassage PassageBrief `json:"toPassageBrief"`
}
