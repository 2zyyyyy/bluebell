package redis

// redis key
const (
	KeyPreFix             = "bluebell:"
	KeyPostTimeZSet       = "post:time"   // 帖子及发帖时间
	KeyPostScoreZSet      = "post:score"  // 帖子及投票的分数
	KeyPostVoteZSetPreFix = "post:voted:" // 记录用户及投票类型;参数是post id
	KeyCommunitySetPreFix = "community:"  // 保存每个分区下的帖子id
)

func getRedisKey(key string) string {
	return KeyPreFix + key
}
