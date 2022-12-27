package application

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"max.workspace.com/cmd/models/errors"
	"max.workspace.com/cmd/models/protocol/application"
)

// initDB db初始化函数
func initDB(config application.DBConfig) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?interpolateParams=true&readTimeout=%vms&writeTimeout=%vms&timeout=%vms&charset=%s",
		config.User, config.Passwd, config.Uri, config.DBName, config.ReadTimeoutMs, config.WriteTimeoutMs, config.TimeoutMs, config.Charset)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(config.ConnMaxOpen)
	db.SetMaxIdleConns(config.ConnMaxIdle)
	return
}

// GetDBHandle 获取db句柄
func GetDBHandle(dbName string) (db *sql.DB, err error) {
	// 检测全局对象是否初始化
	if App == nil {
		err = errors.ErrorApplicationNotInit
		return
	}

	// 检测redis name是否在初始化的列表里
	db, ok := App.DB[dbName]
	if !ok {
		err = errors.ErrorDBNotExist
	}
	return
}
