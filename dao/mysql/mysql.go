package mysql

import (
	"fmt"
	"webapp-scaffold/settings"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		settings.Config.MySQLConfig.User,
		settings.Config.MySQLConfig.Password,
		settings.Config.MySQLConfig.Host,
		settings.Config.MySQLConfig.Port,
		settings.Config.MySQLConfig.Dbname)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(settings.Config.MySQLConfig.MaxOpenCons)
	db.SetMaxOpenConns(settings.Config.MySQLConfig.MaxIdleCons)
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
