package service

import (
	"errors"
	"strconv"
	"webapp-scaffold/dao/mysql"
	"webapp-scaffold/dao/redis"
	"webapp-scaffold/models"

	"go.uber.org/zap"
)

var (
	ErrNotExist = errors.New("当前post id不存在")
)

// CommunityVote 帖子投票功能逻辑处理
func CommunityVote(userID string, p *models.ParamCommunityVote) (err error) {
	// 判断post id对应的帖子是否存在
	var postID = strconv.FormatUint(p.PostID, 10)
	exist, err := mysql.CheckPostExist(postID)
	if err != nil {
		zap.L().Error("mysql.CheckPostExist(post.CommunityID) failed.",
			zap.Uint64("post_id:", p.PostID),
			zap.Error(err))
		return
	}
	if exist == false {
		zap.L().Error("post id not exist.",
			zap.Uint64("post_id:", p.PostID),
			zap.Error(err))
		return ErrNotExist
	}

	// 投票
	err = redis.VoteForCommunity(userID, postID, float64(p.Direction))
	if err != nil {
		zap.L().Debug("CommunityVote",
			zap.String("user_id:", userID),
			zap.Uint64("post_id:", p.PostID),
			zap.Int8("direction:", p.Direction),
			zap.Error(err))
		return
	}
	zap.L().Debug("CommunityVote",
		zap.String("user_id:", userID),
		zap.Uint64("post_id:", p.PostID),
		zap.Int8("direction:", p.Direction))
	return redis.VoteForCommunity(userID, postID, float64(p.Direction))
}
