package show

import (
	"fmt"

	"max.workspace.com/cmd/router"
)

// Cmd 帮助函数，输出可用得子命令
func Cmd() {
	fmt.Println("please refer -cmd cmd!")
	fmt.Println("can use cmd:")
	for cmdName := range router.CmdMap {
		fmt.Println("	" + cmdName)
	}
}
