package models

// User 注册请求参数
type User struct {
	UserID   uint64 `db:"user_id"`
	UserName string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Gender   int    `db:"gender"`
}
