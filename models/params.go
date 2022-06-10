package models

const (
	Page       = 10
	Size       = 1
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
	Email      string `json:"email"`
	Gender     int    `json:"gender"`
}

// ParamLogin 注册请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamOrderList 获取帖子列表
type ParamOrderList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}

// ParamCommunityPostList 社区下帖子列表的接口提
type ParamCommunityPostList struct {
	*ParamOrderList
	CommunityID int64 `json:"community_id" form:"community_id"`
}
