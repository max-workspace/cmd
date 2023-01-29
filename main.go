package main

import (
	"flag"
	"time"

	"max.workspace.com/cmd/application"
	"max.workspace.com/cmd/models/errors"
	app_protocol "max.workspace.com/cmd/models/protocol/application"
	"max.workspace.com/cmd/router"
	"max.workspace.com/cmd/util/show"
)

func main() {
	// 解析命令行参数，找出要执行的指定命令
	cmdParams, err := parseCommandLineParam()
	if err != nil {
		return
	}

	// 解析子命令，生成子命令对象
	cmdGenerator, cmdExist := router.CmdMap[cmdParams.Cmd]
	if !cmdExist {
		show.Cmd()
		err = errors.ErrorCmdNotFind
		return
	}
	cmd := cmdGenerator()

	// 基于命令行参数、子命令配置，初始化全局对象
	application.InitApp(cmdParams, cmd)

	// 子命令执行
	cmd.Exec(cmd)
}

// parseCommandLineParam 解析命令行参数
func parseCommandLineParam() (cmd app_protocol.CmdParams, err error) {
	// 解析命令行参数
	flag.StringVar(&cmd.Cmd, "cmd", "", "executable command")
	flag.StringVar(&cmd.Env, "env", app_protocol.EnvDev, "script execution environment")
	flag.StringVar(&cmd.StartTime, "startTime", time.Now().Local().Format("2006-01-02 15:04:05"), "start time of the script")
	flag.StringVar(&cmd.EndTime, "endTime", time.Now().Local().Format("2006-01-02 15:04:05"), "end time of the script")
	flag.Parse()

	// 检测待执行的命令是否为空
	if len(cmd.Cmd) == 0 {
		show.Cmd()
		err = errors.ErrorCmdNotFind
		return
	}
	return
}
