package service

import (
	"webapp-scaffold/dao/mysql"
	"webapp-scaffold/models"
	"webapp-scaffold/pkg/snowflake"

	"go.uber.org/zap"
)

// 社区相关

// GetCommunityList 处理获取社区列表逻辑
func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

// GetCommunityDetail 处理获取社区列表逻辑
func GetCommunityDetail(id int64) (detail *models.CommunityDetail, err error) {
	return mysql.GetCommunityByID(id)
}

func CreateCommunityPost(p *models.CommunityPost) (err error) {
	// 1.生成id
	id, err := snowflake.GenID()
	if err != nil {
		zap.L().Error("post snowflake.GenID failed.", zap.Error(err))
		return
	}
	// 2.保存到数据库

	// 3.返回响应
	return
}
