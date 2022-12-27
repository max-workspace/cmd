package application

const (
	// DBNameTest DB配置名称test
	DBNameTest = "test"

	// RedisNameTest Redis配置名称test
	RedisNameTest = "test"
)

// Conf 全局配置

type Config struct {
	Logger LoggerConfig           `yaml:"logger"`
	DB     map[string]DBConfig    `yaml:"db"`
	Redis  map[string]RedisConfig `yaml:"redis"`
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	PathDir      string `yaml:"path_dir"`
	Level        int    `yaml:"level"`
	ReportCaller bool   `yaml:"report_caller"`
}

// DBConfig db配置
type DBConfig struct {
	Uri            string `yaml:"uri"`
	User           string `yaml:"user"`
	Passwd         string `yaml:"passwd"`
	DBName         string `yaml:"db_name"`
	ReadTimeoutMs  string `yaml:"read_timeout_ms"`
	WriteTimeoutMs string `yaml:"write_timeout_ms"`
	TimeoutMs      string `yaml:"timeout_ms"`
	Charset        string `yaml:"charset"`
	ConnMaxOpen    int    `yaml:"conn_max_open"`
	ConnMaxIdle    int    `yaml:"conn_max_idle"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Uri            string `yaml:"uri"`
	MaxIdle        int    `yaml:"max_idle"`
	MaxActive      int    `yaml:"max_active"`
	IdleTimeOutS   int    `yaml:"idle_time_out_s"`
	Wait           bool   `yaml:"wait"`
	ConnTimeoutMs  int    `yaml:"conn_timeout_ms"`
	ReadTimeoutMs  int    `yaml:"read_timeout_ms"`
	WriteTimeoutMs int    `yaml:"write_timeout_ms"`
}
