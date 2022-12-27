package router

import (
	"max.workspace.com/cmd/cmd/hello_world"
	"max.workspace.com/cmd/models/protocol/cmd"
)

var (
	// CmdMap 命令行命令列表
	CmdMap = map[string]cmd.CmdGenerator{
		"hello_world": cmd.CmdGenerator(hello_world.NewCmd),
	}
)
