package application

const (
	// EnvDev 环境标识变量
	EnvDev = "dev"
	// EnvOnline 环境标识变量
	EnvOnline = "online"

	// EnvOnlineSuffix 环境变量后缀
	EnvOnlineSuffix = ".online"
)

var (
	// ConfigPathSuffixMap 配置文件在不同环境的后缀
	ConfigPathSuffixMap = map[string]string{
		EnvDev:    "",
		EnvOnline: EnvOnlineSuffix,
	}
)

// CmdParams 命令行参数
type CmdParams struct {
	Cmd       string `json:"cmd"`
	Env       string `jsob:"env"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
