package models

type ParamCommunityVote struct {
	PostID    uint64 `json:"post_id,string" binding:"required"`                // 帖子id
	Direction int8   `json:"direction,string" binding:"required,oneof=1 0 -1"` // 赞成：1 反对：-1 取消：0
}
