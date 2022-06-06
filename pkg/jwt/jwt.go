package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenExpireDuration token过期时间
const TokenExpireDuration = time.Hour * 72

var (
	mySecret   = []byte("设置你需要的秘钥")
	ErrorToken = errors.New("无效的token")
)

type MyClaims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID uint64, username string) (string, error) {
	// 创建一个我们自己声明的数据
	myClaims := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "2zyyyyy",                                  // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	// 使用指定的secret签名并获得完成的编码后的字符串
	return token.SignedString(mySecret)
}

// ParseToken 解析JWT
func ParseToken(token string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims)
	t, err := jwt.ParseWithClaims(token, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if t.Valid { // 校验token
		return mc, err
	}
	return nil, ErrorToken
}
