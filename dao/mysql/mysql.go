package mysql

import (
	"fmt"
	"web_app/settings"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(config *settings.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Dbname,
	)
	db, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		fmt.Printf("数据库连接错误,err:%v\n", err)
		zap.L().Error("database connect failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	fmt.Printf("数据库连接成功!\n")
	zap.L().Info("database connect success")
	return
}

func Close() {
	_ = db.Close()
	return
}
