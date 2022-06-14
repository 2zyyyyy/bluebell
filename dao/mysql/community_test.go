package mysql

import (
	"bluebell/models"
	"bluebell/settings"
	"testing"
)

func init() {
	config := &settings.MySQLConfig{
		Host:        "127.0.0.1",
		Port:        3306,
		User:        "root",
		Password:    "123456",
		Dbname:      "bluebell",
		MaxOpenCons: 200,
		MaxIdleCons: 10,
	}
	err := Init(config)
	if err != nil {
		panic(err)
	}
}

func TestCreateCommunityPost(t *testing.T) {
	communityPost := &models.CommunityPost{
		ID:          1,
		AuthorID:    10,
		CommunityID: 1,
		Title:       "dao.mysql 的单元测试",
		Content:     "创建帖子函数单元测试用例",
	}
	err := CreateCommunityPost(communityPost)
	if err != nil {
		t.Fatalf("CreateCommunityPost insert into record mysql failed, err%v\n", err)
	}
	t.Logf("CreateCommunityPost insert record into mysql success.")
}
