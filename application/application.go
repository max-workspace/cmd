package application

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/gomodule/redigo/redis"
	yaml "gopkg.in/yaml.v2"

	"max.workspace.com/cmd/models/errors"
	"max.workspace.com/cmd/models/protocol/application"
	"max.workspace.com/cmd/models/protocol/cmd"
)

const (
	ConfigPath = "config/cmd.yaml"
)

var (
	// 应用单例
	app *Application

	// 保证整体服务实例只启动一次
	onceApp sync.Once
)

// Application 全局应用
type Application struct {
	// CmdParams 命令行参数
	CmdParams application.CmdParams `json:"cmd_params"`

	// ConfigPath 实例化后的配置路径
	ConfigPath string `json:"config_path"`

	// redis 全局实例
	Redis map[string]*redis.Pool `json:"redis"`

	// db 全局实例
	DB map[string]*sql.DB `json:"db"`
}

// InitApp 初始化应用实例
func InitApp(cmdParams application.CmdParams, cmd cmd.Cmd) *Application {
	onceApp.Do(func() {
		// 初始化应用
		app = new(Application)
		app.CmdParams = cmdParams
		app.ConfigPath = cmd.GetConfigPath()
		app.Redis = make(map[string]*redis.Pool)
		app.DB = make(map[string]*sql.DB)

		// 基于环境信息以及基础配置地址生成最终配置加载地址
		realPath, err := GetConfigPath(app.ConfigPath, cmdParams)
		if err != nil {
			panic(fmt.Sprintf("GetConfigPath fail! err=[%v]\n", err))
		}

		// 根据路径加载基础配置
		yamlFile, err := os.ReadFile(realPath)
		if err != nil {
			panic(fmt.Sprintf("read config file fail! err=[%v]\n", err))
		}

		// 解析配置文件
		cmdConfig := new(application.Config)
		err = yaml.Unmarshal(yamlFile, &cmdConfig)
		if err != nil {
			panic(fmt.Sprintf("parse config fail! err=[%v]\n", err))
		}

		// 初始化logrus
		initLog(cmdParams, cmdConfig.Logger)

		// 基于命令配置选择需要初始化redis
		redisList := cmd.GetRedisList()
		for _, redisName := range redisList {
			redisConfig, ok := cmdConfig.Redis[redisName]
			if !ok {
				panic(fmt.Sprintf("redis config not find! redisName=[%v]\n", redisName))
			}
			redisPool, err := initRedis(redisConfig)
			if err != nil {
				panic(fmt.Sprintf("redis init fail! redisConfig=[%v] err=[%v]\n", redisConfig, err))
			}
			app.Redis[redisName] = redisPool
		}

		// 基于命令配置选择需要初始化mysql
		dbList := cmd.GetDBList()
		for _, dbName := range dbList {
			dbConfig, ok := cmdConfig.DB[dbName]
			if !ok {
				panic(fmt.Sprintf("db config not find! dbName=[%v]\n", dbName))
			}
			db, err := initDB(dbConfig)
			if err != nil {
				panic(fmt.Sprintf("db init fail! dbConfig=[%v] err=[%v]\n", dbConfig, err))
			}
			app.DB[dbName] = db
		}
	})
	return app
}

// NewApp 获取应用实例
func NewApp() *Application {
	return app
}

// GetConfigPath 根据环境信息以及基础配置路径
func GetConfigPath(configpath string, cmdParams application.CmdParams) (realPath string, err error) {
	suffix, ok := application.ConfigPathSuffixMap[cmdParams.Env]
	if !ok {
		err = errors.ErrorConfigEnvNotFind
		return
	}
	realPath = configpath + suffix
	return
}
