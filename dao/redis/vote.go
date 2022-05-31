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
	oneWeekSeconds = 70 * 24 * 3600 // 帖子投票过期时间
	oneTicketScore = 432            // 每票的分数
)

var ErrVoteExpire = errors.New("吉时已过")

func VoteForCommunity(userID, postID uint64, value float64) (err error) {
	// 1.判断投票限制
	// 从redis获取帖子发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), strconv.FormatUint(postID, 10)).Val()
	if float64(time.Now().Unix())-postTime > oneWeekSeconds {
		return ErrVoteExpire
	}
	//2.更新帖子分数
	//2.1 查询当前用户给当前帖子的投票记录
	oldValue := rdb.ZScore(getRedisKey(KeyPostVoteZSetPreFix+strconv.FormatUint(postID, 10)),
		strconv.FormatUint(userID, 10)).Val()
	var op float64
	if value > oldValue { // 如果当前分数大于历史分数
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(value - oldValue) // 计算两次投票的差值
	rdb.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*oneTicketScore, strconv.FormatUint(postID, 10))

	//3.记录用户为该帖子投票的数据
	if value == 0 {
		// 取消投票
		_, err = rdb.ZRem(getRedisKey(KeyPostVoteZSetPreFix+strconv.FormatUint(postID, 10)), postID).Result()
	} else {
		_, err = rdb.ZAdd(getRedisKey(KeyPostVoteZSetPreFix+strconv.FormatUint(postID, 10)), redis.Z{
			Score:  value, // 赞成/反对
			Member: userID,
		}).Result()
	}
	return
}
