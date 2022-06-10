package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

/*	投票功能场景分析
	case 1 direction = 1:
		1.1 用户没有投过票 投了赞成票	+1*432
		1.2 用户投过反对票 改投赞成票  	+2*432
	case 2 direction = 0:
		2.1 用户投过赞成票 取消了投票	-1*432
		2.2 用户投过反对票 取消了投票	+1*432
	case 3 direction = -1:
		3.1 用户没有投过票 投了赞成票	+1*432
		3.2 用户投过赞成票 改投反对票	-2*432
	投票功能限制
	每个帖子自发布之日起7日内允许投票 超过改时间则不允许再投票
	1.到期之后将redis中保存的帖子对应赞成和反对票存储至mysql中
	2.到期之后删除 KeyPostVoteZSetPreFix */

const (
	oneWeekSeconds         = 7 * 24 * 3600 // 帖子投票过期时间
	oneTicketScore float64 = 5             // 每票的分数
)

var (
	ErrVoteExpire     = errors.New("吉时已过")
	ErrVoteRepetition = errors.New("不允许重复投票")
)

// CreateCommunityPost 创建帖子存储时间
func CreateCommunityPost(postID, communityID int64) (err error) {
	// 使用事务更新redis数据
	pipeline := rdb.TxPipeline()
	// 更新帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 更新帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	// 把帖子id加入到社区的set中
	communityKey := getRedisKey(KeyCommunitySetPreFix + strconv.FormatInt(communityID, 10))
	pipeline.SAdd(communityKey, postID)
	_, err = pipeline.Exec()
	return
}

// VoteForCommunity 投票功能函数
func VoteForCommunity(userID, postID string, value float64) (err error) {
	// 1.判断投票限制
	// 从redis获取帖子发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekSeconds {
		return ErrVoteExpire
	}
	//2.更新帖子分数
	//2.1 查询当前用户给当前帖子的投票记录
	oldValue := rdb.ZScore(getRedisKey(KeyPostVoteZSetPreFix+postID), userID).Val()
	var op float64
	// 校验是否重复投票
	if value == oldValue {
		return ErrVoteRepetition
	}
	if value > oldValue { // 如果当前分数大于历史分数
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(oldValue - value) // 计算两次投票的差值
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*oneTicketScore, postID)

	//3.记录用户为该帖子投票的数据
	if value == 0 {
		// 取消投票
		pipeline.ZRem(getRedisKey(KeyPostVoteZSetPreFix+postID), postID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVoteZSetPreFix+postID), redis.Z{
			Score:  value, // 赞成/反对
			Member: userID,
		})
	}
	_, err = pipeline.Exec()
	return
}
