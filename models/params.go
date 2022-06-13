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
	Page        int64  `json:"page" form:"page"`                 // 页码
	Size        int64  `json:"size" form:"size"`                 // 每页数据量
	Order       string `json:"order" form:"order"`               // 排序依据
	CommunityID int64  `json:"community_id" form:"community_id"` // 可以为空
}

// ParamCommunityPostList 社区下帖子列表的接口提
//type ParamCommunityPostList struct {
//	*ParamOrderList
//	CommunityID int64 `json:"community_id" form:"community_id"`
//}
