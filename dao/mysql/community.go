package mysql

import (
	"database/sql"
	"webapp-scaffold/models"

	"go.uber.org/zap"
)

// 社区相关

// GetCommunityList 查询所有社区
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		// 如果查询为空
		if err == sql.ErrNoRows {
			zap.L().Warn("no result from community table")
			err = nil
		}
	}
	return
}

// GetCommunityByID 根据社区ID查询社区详情
func GetCommunityByID(id int64) (communityDetail *models.CommunityDetail, err error) {
	// 申请内存
	communityDetail = new(models.CommunityDetail)
	sqlStr := "select community_id, community_name, introduction, create_time from community where id = ?"
	if err := db.Get(communityDetail, sqlStr, id); err != nil {
		// 判断id是否有效
		if err == sql.ErrNoRows {
			//controllers.ResponseError(c, ErrorInvalidID)
			err = ErrorInvalidID
			return nil, err
		}
	}
	return
}
