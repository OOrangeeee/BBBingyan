package infoModels

import "time"

type Passage struct {
	ID                    uint   `json:"id"`
	PassageTitle          string `json:"passage_title"`
	PassageContent        string `json:"passage_content"`
	PassageAuthorUserName string `json:"passage_author_user_name"`
	PassageAuthorNickName string `json:"passage_author_nick_name"`
	PassageTag            string `json:"passage_tag"`
	PassageTime           time.Time
}
