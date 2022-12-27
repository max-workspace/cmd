package hello_world

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"max.workspace.com/cmd/application"
	app_protocol "max.workspace.com/cmd/models/protocol/application"
	helloworld "max.workspace.com/cmd/service/page/hello_world"
)

// Run 脚本命令实际需要执行的方法
func (c *CMD) Run() {
	app := application.NewApp()
	log.Info(fmt.Sprintf("cmd=[%+v] ExtraConfig=[%+v]", app.CmdParams.Cmd, c.ExtraConfig))

	// 展示redis使用
	helloworldPageService := helloworld.New()
	ret, err := helloworldPageService.GetRedisHelloWorld()
	log.Info(fmt.Sprintf("GetRedisHelloWorld! err=[%+v] ret=[%+v]", err, ret))
	if err != nil {
		log.Warn(fmt.Sprintf("GetRedisHelloWorld fail!! err=[%+v] ret=[%+v]", err, ret))
	}

	// 展示db使用
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	dbHandle, err := application.GetDBHandle(app_protocol.DBNameTest)
	if err != nil {
		fmt.Printf("GetDBHandle fail! err:%+v\n", err)
	}
	sql := fmt.Sprintf("select table_name from information_schema.tables limit 2;")
	rows, err := dbHandle.QueryContext(ctx, sql)
	if err != nil {
		fmt.Printf("QueryContext fail! sql:%+v err:%+v\n", sql, err)
	}
	tableNameList := make([]string, 0)
	for rows.Next() {
		var tableName string
		err = rows.Scan(
			&tableName,
		)
		if err != nil {
			fmt.Printf("rows.Scan fail! err:%+v\n", err)
		}

		tableNameList = append(tableNameList, tableName)
	}
	fmt.Printf("tableNameList:%+v\n", tableNameList)

	log.Info(fmt.Sprintf("cmd:%v end", app.CmdParams.Cmd))
}
