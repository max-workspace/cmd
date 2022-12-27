package cmd

import (
	"time"

	"max.workspace.com/cmd/application"
	"max.workspace.com/cmd/models/protocol/cmd"
)

type BaseCMD struct {
	RunStartTime time.Time
	RunEndTime   time.Time
}

// GetConfigPath 命令行命令注册所需接口实现，默认全局配置的加载路径，允许子类进行重写制定全局配置文件
func (c *BaseCMD) GetConfigPath() string {
	return application.ConfigPath
}

// Init 命令行命令注册所需接口实现 初始化
func (c *BaseCMD) Init(cmd cmd.Cmd) {
	cmd.ExtraConfigInit()
}

// RunBefore 命令行命令注册所需接口实现 允许子类进行重写定制脚本实际执行前的操作
func (c *BaseCMD) RunBefore() {
	c.RunStartTime = time.Now().Local()
}

// RunBefore 命令行命令注册所需接口实现 允许子类进行重写定制脚本实际执行后的操作
func (c *BaseCMD) RunAfter() {
	c.RunEndTime = time.Now().Local()
}
