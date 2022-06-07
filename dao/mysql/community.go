package mysql

import (
	"database/sql"
	"strings"
	"webapp-scaffold/models"

	"github.com/jmoiron/sqlx"

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
	sqlStr := "select post_id, title, content, author_id, community_id, status, create_time from post order by create_time desc limit ?,?"
	err = db.Select(&list, sqlStr, (page-1)*size, size)
	return
}

// CheckPostExist 检查帖子是否存在
func CheckPostExist(id string) (exist bool, err error) {
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

// GetPostOrderList 根据redis查询的id查询对应的帖子详情
func GetPostOrderList(ids []string) (postList []*models.CommunityPost, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
				from post
				where post_id in (?) 
				order by FIND_IN_SET(post_id, ?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))

	/*
		FIND_IN_SET(str,strList)

		str 要查询的字符串
		strList 字段名，参数以“,”分隔，如(1,2,6,8)
		查询字段(strList)中包含的结果，返回结果null或记录。
		strList 中，则返回值的范围在 1 到 N 之间。
		一个字符串列表就是一个由一些被 ‘,’ 符号分开的子链组成的字符串。如果第一个参数是一个常数字符串，
		而第二个是type SET列，则FIND_IN_SET() 函数被优化，使用比特计算。
		strList strList 为空字符串，则返回值为 0 。
		如任意一个参数为NULL，则返回值为 NULL。这个函数在第一个参数包含一个逗号(‘,’)时将无法正常运行。
	*/
	if err != nil {
		return nil, err
	}
	// sqlx.In 返回带 `?` bind-var的查询语句, 我们使用Rebind()重新绑定它
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
