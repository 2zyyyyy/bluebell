package redis

import (
	"strconv"
	"time"
	"webapp-scaffold/models"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	// 1.确定查询的索引起始点
	start := (page - 1) * size
	end := start + size - 1
	// 2.zrevrange 查询(按分数从大到小的顺序查询指定数量的元素)
	return rdb.ZRevRange(key, start, end).Result()
}

// GetPostListByID 根据id查询帖子列表
func GetPostListByID(p *models.ParamOrderList) ([]string, error) {
	// 从redis获取id
	// 1.根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return getIDsFormKey(key, p.Page, p.Size)
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

// GetCommunityPostListByID 根据社区id返回对应的帖子id列表
func GetCommunityPostListByID(p *models.ParamOrderList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}

	// 1.使用ZInterStore把分区的帖子set和帖子分数的ZSet生成一个新的ZSet
	// 2.针对新的ZSet按之前的逻辑去数据

	// 社区的key
	communityKey := getRedisKey(KeyCommunitySetPreFix) + strconv.FormatInt(p.CommunityID, 10)
	// 3.使用缓存key减少ZInterStore执行的次数
	key := orderKey + strconv.FormatInt(p.CommunityID, 10)
	if rdb.Exists(key).Val() < 1 {
		// 不存在 需要计算
		pipeline := rdb.TxPipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, communityKey, orderKey)
		// 设置超时时间
		pipeline.Expire(key, time.Second*60)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	// 如果存在就直接根据key查询对应的ids
	return getIDsFormKey(key, p.Page, p.Size)
}
