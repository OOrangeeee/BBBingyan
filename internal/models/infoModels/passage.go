package infoModels

import "time"

type PassageDetail struct {
	ID                    uint      `json:"id"`
	PassageTitle          string    `json:"passageTitle"`
	PassageContent        string    `json:"passageContent"`
	PassageAuthorUserName string    `json:"passageAuthorUserName"`
	PassageAuthorNickName string    `json:"passageAuthorNickName"`
	PassageAuthorId       uint      `json:"passageAuthorId"`
	PassageTag            string    `json:"passageTag"`
	PassageBeLikedCount   int       `json:"passageBeLikedCount"`
	PassageCommentCount   int       `json:"passageCommentCount"`
	PassageTime           time.Time `json:"passageSendTime"`
}

type PassageBrief struct {
	ID           uint   `json:"id"`
	PassageTitle string `json:"passageTitle"`
	// no PassageContent
	PassageAuthorUserName string `json:"passageAuthorUserName"`
	PassageAuthorNickName string `json:"passageAuthorNickName"`
	// no PassageAuthorId
	PassageTag          string    `json:"passageTag"`
	PassageBeLikedCount int       `json:"passageBeLikedCount"`
	PassageCommentCount int       `json:"passageCommentCount"`
	PassageTime         time.Time `json:"passageSendTime"`
}
