package controllers

import "bluebell/models"

// 帖子详情接口响应数据
type _ResponsePostDetail struct {
	Code    ResCode                 `json:"code"`    // 状态码
	Message string                  `json:"message"` // 提示信息
	Data    *models.CommunityDetail `json:"data"`    // 数据
}
