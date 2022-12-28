package hello_world

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"max.workspace.com/cmd/cmd"
	"max.workspace.com/cmd/models/protocol/application"
	cmd_protocol "max.workspace.com/cmd/models/protocol/cmd"
)

var (
	// DBList 需要使用到的DB配置列表
	DBList = []string{
		application.DBNameTest,
	}

	// RedisList 需要使用到的Redis配置列表
	RedisList = []string{
		application.RedisNameTest,
	}
)

const (
	// ExtraConfigPath 自定义的扩展配置路径，按照具体需求选择进行添加
	ExtraConfigPath = "cmd/hello_world/config/hello_world.yaml"
)

// CMD 子命令对象
type CMD struct {
	cmd.BaseCMD

	// 自定义的扩展配置，按照具体需求选择进行添加
	ExtraConfig ExtraInfoConfig `json:"extra_info_config"`
}

// ExtraInfoConfig 当前脚本加载的额外配置
type ExtraInfoConfig struct {
	Desc       string         `yaml:"desc"`
	StatusList []int          `yaml:"status_list"`
	StatusMap  map[string]int `yaml:"status_map"`
}

// NewCmd 路由需要注册的函数值
func NewCmd() cmd_protocol.Cmd {
	return &CMD{}
}

// GetDBList 全局变量初始化需要的DB连接池的配置名称
func (c *CMD) GetDBList() []string {
	return DBList
}

// GetRedisList 全局变量初始化需要的redis连接池的配置名称
func (c *CMD) GetRedisList() []string {
	return RedisList
}

// ExtraConfigInit 命令行命令注册所需接口实现，按照自身需求进行实现
func (c *CMD) ExtraConfigInit() {
	// 读取配置文件
	yamlFile, err := os.ReadFile(ExtraConfigPath)
	if err != nil {
		panic(fmt.Sprintf("ExtraConfigInit ReadFile fail! ExtraConfigPath=[%v] err=[%v]\n", ExtraConfigPath, err))
	}

	// 解析配置文件
	var extraConfig ExtraInfoConfig
	err = yaml.Unmarshal(yamlFile, &extraConfig)
	if err != nil {
		panic(fmt.Sprintf("ExtraConfigInit Unmarshal fail! yamlFile=[%s] err=[%v]\n", string(yamlFile), err))
	}
	c.ExtraConfig = extraConfig
}
