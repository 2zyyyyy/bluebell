package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestCreatePostHandler(t *testing.T) {
	// 设置gin模式
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// 定义测试请求接口路径
	url := "/community/post"
	// 组装请求
	r.POST(url, CreatePostHandler)

	// 测试数据
	body := `{"community_id": 1, "title": "Test", "content": "测试创建帖子接口的测试用例"}`
	// 发起请求
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 断言
	assert.Equal(t, 200, w.Code)
	// 判断相应是否返回需要登录
	assert.Contains(t, w.Body.String(), "请先登录")
}
