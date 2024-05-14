package infoModels

type Comment struct {
	CommentContent string       `json:"commentContent"`
	FromUser       User         `json:"fromUser"`
	ToPassage      PassageBrief `json:"toPassageBrief"`
}
