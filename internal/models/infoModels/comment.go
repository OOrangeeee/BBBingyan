package infoModels

type Comment struct {
	CommentContent string  `json:"commentContent"`
	FromUser       User    `json:"fromUser"`
	ToPassage      Passage `json:"toPassage"`
}
