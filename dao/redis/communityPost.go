package redis

import (
	"webapp-scaffold/models"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// GetPostListByID 根据id查询帖子列表
func GetPostListByID(p *models.ParamOrderList) ([]string, error) {
	// 从redis获取id
	// 1.根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 2.确定查询的索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	// 3.zrevrange 查询
	return rdb.ZRevRange(key, start, end).Result()
}

// GetPostVoteData 根据帖子的id去redis查询对应赞成票的数量
func GetPostVoteData(ids []string) (data []int64, err error) {
	// 初始化切片
	//data = make([]int64, 0, len(ids))
	//// 遍历ids获取id的投票数
	//for _, id := range ids {
	//	// 获取id对应redis的key值
	//	key := getRedisKey(KeyPostVoteZSetPreFix + id)
	//	// 根据key查询对应的数据
	//	v := rdb.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}

	// 使用pipeline一次发送多条命令 减少RTT
	pipeline := rdb.TxPipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVoteZSetPreFix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	zap.L().Debug("cmders:", zap.Any("cmders", cmders))
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	zap.L().Debug("data:", zap.Any("data", data))
	return
}
