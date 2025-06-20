// 启动程序
package main

import (
	"blog/core"
	"blog/global"
	"fmt"
)

func main() {
	//初始化logger
	global.Log = core.InitLogger()
	// core.MysqlInit()
	fmt.Println(core.RedisInit())

}
