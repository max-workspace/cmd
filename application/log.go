package application

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"max.workspace.com/cmd/models/protocol/application"
	"max.workspace.com/cmd/util/file"
)

var (
	// LevelMap 日志等级映射配置
	LevelMap = map[int]log.Level{
		0: log.PanicLevel,
		1: log.FatalLevel,
		2: log.ErrorLevel,
		3: log.WarnLevel,
		4: log.InfoLevel,
		5: log.DebugLevel,
		6: log.TraceLevel,
	}

	// LogAlreadyInit 日志初始化状态
	LogAlreadyInit = false
)

// initLog log初始化函数
func initLog(cmdParams application.CmdParams, config application.LoggerConfig) {
	// 检测日志级别映射是否存在
	level, ok := LevelMap[config.Level]
	if !ok {
		panic(fmt.Sprintf("level map not find! level:%v\n", config.Level))
	}

	// 设置日志记录等级
	log.SetLevel(level)
	// 日志中添加文件名和方法信息
	log.SetReportCaller(config.ReportCaller)
	// 设置日志格式为json
	log.SetFormatter(&log.JSONFormatter{})

	// 检测日志路径是否存在
	exist, err := file.PathExists(config.PathDir)
	if err != nil {
		panic(fmt.Sprintf("PathExists fail! path:%v err:%+v\n", config.PathDir, err))
	}
	// 如果不存在创建对应目录
	if !exist {
		err := os.Mkdir(config.PathDir, 0666)
		if err != nil {
			panic(fmt.Sprintf("create log dir fail! path:%v err:%+v\n", config.PathDir, err))
		}
	}

	// 生成日志文件名称
	logFile := config.PathDir + cmdParams.Cmd + ".log"
	// 设置日志持久化的信息
	writer, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(fmt.Sprintf("OpenFile fail! err=[%v]\n", err))
	}
	log.SetOutput(writer)
	LogAlreadyInit = true
}
