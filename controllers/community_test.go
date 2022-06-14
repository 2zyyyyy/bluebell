package controllers

import (
	"bytes"
	"encoding/json"
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
	// 方法一：判断响应内容中是不是包含指定的字符串
	assert.Contains(t, w.Body.String(), "请先登录")

	// 方法二：将响应的数据反序列化至ResponseData 然后判断字段值是否符合预期结果
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json unmarshal w.body failed. err%s", err)
	}
	assert.Equal(t, res.Code, CodeNeedLogin)
}
