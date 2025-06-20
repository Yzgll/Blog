// 启动程序
package main

import (
	"blog/core"
	"blog/global"
)

func main() {
	//初始化logger
	global.Log = core.InitLogger()
	core.MysqlInit()
}
