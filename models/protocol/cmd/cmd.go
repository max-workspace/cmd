package cmd

type Cmd interface {
	// 需要业务自身实现的
	Run()
	GetDBList() []string
	GetRedisList() []string
	ExtraConfigInit()

	// 基类已经实现的，允许子类重写的方法
	Exec(cmd Cmd)
	Init(cmd Cmd)
	GetConfigPath() string
	RunBefore()
	RunAfter()
}

// CmdGenerator
type CmdGenerator func() Cmd
