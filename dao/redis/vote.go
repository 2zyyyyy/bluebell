package redis

import (
	"errors"
	"strconv"
	"time"
)

const oneWeekSeconds = 7 * 24 * 3600

var ErrVoteExpire = errors.New("吉时已过")

func VoteForCommunity(userID, postID uint64, value float64) (err error) {
	// 1.判断投票限制
	// 从redis获取帖子发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), strconv.FormatUint(postID, 10)).Val()
	if float64(time.Now().Unix())-postTime > oneWeekSeconds {
		return ErrVoteExpire
	}
	//2.更新帖子分数

	//3.记录用户为该帖子投票的数据
	return
}
