package router

import (
	"blog/api"
	"blog/config"

	"github.com/gin-gonic/gin"
)

//初始化路由

func InitRouter() *gin.Engine {
	gin.SetMode(config.Config.System.Env)
	router := gin.New()
	router.Use(gin.Recovery())
	register(router)
	return router
}

func register(r *gin.Engine) {
	r.GET("/api/success", api.Success)
	r.GET("/api/failed", api.Failed)
}
