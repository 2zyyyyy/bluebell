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

// CreateCommunityPost 创建社区的帖子
func CreateCommunityPost(post *models.CommunityPost) (err error) {
	sqlStr := "insert into post(post_id, title, author_id, community_id, content) value(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, post.ID, post.Title, post.AuthorID, post.CommunityID, post.Content)
	if err != nil {
		return err
	}
	return
}

// GetAuthorNameById 根据用户id查询用户名称
func GetAuthorNameById(userId uint64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := "select user_id, username from user where user_id = ?"
	if err := db.Get(user, sqlStr, userId); err != nil {
		// 判断id是否有效
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
			return nil, err
		}
	}
	return
}

// GetPostDetailByID 根据帖子ID查询帖子详情
func GetPostDetailByID(postId uint64) (postDetail *models.CommunityPost, err error) {
	// 申请内存
	postDetail = new(models.CommunityPost)
	sqlStr := "select post_id, title, content, author_id, community_id, status, create_time from post where post_id = ?"
	if err := db.Get(postDetail, sqlStr, postId); err != nil {
		// 判断id是否有效
		if err == sql.ErrNoRows {
			//controllers.ResponseError(c, ErrorInvalidID)
			err = ErrorInvalidID
			return nil, err
		}
	}
	return
}

// GetPostList 查询帖子列表
func GetPostList(page, size int64) (list []*models.CommunityPost, err error) {
	list = make([]*models.CommunityPost, 0, 2)
	sqlStr := "select post_id, title, content, author_id, community_id, status, create_time from post limit ?,?"
	err = db.Select(&list, sqlStr, (page-1)*size, size)
	return
}

// CheckPostExist 检查帖子是否存在
func CheckPostExist(id uint64) (exist bool, err error) {
	var count int
	sqlStr := "select count(community_id) from post where post_id = ?"
	if err := db.Get(&count, sqlStr, id); err != nil {
		return exist, err
	}
	zap.L().Debug("select count from post", zap.Int("count:", count))
	if count > 0 {
		return true, nil
	}
	return exist, nil
}
