package service

import (
	"webapp-scaffold/dao/mysql"
	"webapp-scaffold/dao/redis"
	"webapp-scaffold/models"
	"webapp-scaffold/pkg/snowflake"

	"go.uber.org/zap"
)

// 社区相关

// GetCommunityList 获取社区列表逻辑
func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

// GetCommunityDetail 获取社区列表逻辑
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityByID(id)
}

// CreateCommunityPost 获取帖子详情逻辑
func CreateCommunityPost(post *models.CommunityPost) (err error) {
	// 1.生成id
	var id uint64
	id, err = snowflake.GenID()
	if err != nil {
		zap.L().Error("create post snowflake.GenID failed.",
			zap.Error(err))
		return
	}
	post.ID = int64(id)
	// 2.保存到数据库并返回
	err = mysql.CreateCommunityPost(post)
	if err != nil {
		return err
	}
	// 保存到redis
	err = redis.CreateCommunityPost(int64(id))
	return
}

// GetPostDetail 获取帖子详情逻辑
func GetPostDetail(id uint64) (apiPostDetail *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostDetailByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID(id) failed.",
			zap.Uint64("authorID:", id),
			zap.Error(err))
		return
	}
	// 1.根据作者id查询作者用户名
	author, err := mysql.GetAuthorNameById(uint64(post.AuthorID))
	if err != nil {
		zap.L().Error("mysql.GetAuthorNameById(post.AuthorID) failed.",
			zap.Int64("authorID:", post.AuthorID),
			zap.Error(err))
		return
	}
	// 2.根据社区id查询社区名称
	community, err := mysql.GetCommunityByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityNameById(post.CommunityID) failed.",
			zap.Int64("authorID:", post.CommunityID),
			zap.Error(err))
		return
	}
	apiPostDetail = &models.ApiPostDetail{
		AuthorName:      author.UserName,
		CommunityDetail: community,
		CommunityPost:   post,
	}
	return
}

// GetPostList 获取帖子列表逻辑
func GetPostList(page, size int64) (list []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList failed.", zap.Error(err))
		return
	}
	list = make([]*models.ApiPostDetail, 0, len(posts))
	// 循环posts获取用户名和社区名称
	for _, post := range posts {
		// 1.根据作者id查询作者用户名
		author, err := mysql.GetAuthorNameById(uint64(post.AuthorID))
		if err != nil {
			zap.L().Error("mysql.GetAuthorNameById(post.AuthorID) failed.",
				zap.Int64("authorID:", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 1.根据社区id查询社区名称
		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetAuthorNameById(post.AuthorID) failed.",
				zap.Int64("authorID:", post.AuthorID),
				zap.Error(err))
			continue
		}
		apiPostDetail := &models.ApiPostDetail{
			AuthorName:      author.UserName,
			CommunityDetail: community,
			CommunityPost:   post,
		}
		list = append(list, apiPostDetail)
	}
	return
}
