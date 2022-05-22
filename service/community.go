package service

import (
	"webapp-scaffold/dao/mysql"
	"webapp-scaffold/models"
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
