package infoModels

type Comment struct {
	CommentContent string  `josn:"commentContent"`
	FromUser       User    `json:"fromUser"`
	ToPassage      Passage `json:"toPassage"`
}
