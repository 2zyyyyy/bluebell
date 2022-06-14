package mysql

import (
	"bluebell/settings"
	"fmt"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(config *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Dbname)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(config.MaxOpenCons)
	db.SetMaxOpenConns(config.MaxIdleCons)
	return
}

// Close 封装db.close方法
func Close() {
	err := db.Close()
	if err != nil {
		zap.L().Error("mysql close failed.", zap.Error(err))
		return
	}
}
