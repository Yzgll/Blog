// 启动程序
package main

import (
	"blog/config"
	"blog/core"
	"blog/global"
	"blog/router"
	"fmt"
)

func main() {
	//初始化logger
	global.Log = core.InitLogger()
	fmt.Println(core.MysqlInit())
	fmt.Println(core.RedisInit())
	//初始化路由
	router := router.InitRouter()
	address := fmt.Sprintf("%s:%d", config.Config.System.Host, config.Config.System.Port)
	global.Log.Infof("系统启动成功，运行在：%s", address)
	router.Run(address)
}
