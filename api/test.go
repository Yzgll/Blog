package api

import (
	"blog/result"

	"github.com/gin-gonic/gin"
)

// 测试相关接口
// 成功测试
func Success(c *gin.Context) {
	result.Sucess(c, 200)
}

func Failed(c *gin.Context) {
	result.Failed(c, int(result.ApiCode.Failed), result.ApiCode.GetMessage(result.ApiCode.Failed))
}
