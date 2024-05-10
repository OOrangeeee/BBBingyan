package infoModels

type User struct {
	UserId           uint   `json:"userId"`
	UserName         string `json:"userName"`
	UserEmail        string `json:"userEmail"`
	UserNickName     string `json:"userNickName"`
	UserFollowCount  int    `json:"userFollowCount"`
	UserFansCount    int    `json:"userFansCount"`
	UserPassageCount int    `json:"userPassageCount"`
	UserLikeCount    int    `json:"userLikeCount"`
	UserIsAdmin      bool   `json:"userIsAdmin"`
}
