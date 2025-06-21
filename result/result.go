package result

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//定义返回的结构体

type Result struct {
	Code    int         `json:"code"`    //状态码
	Message string      `json:"message"` //提示信息
	Data    interface{} `json:"data"`    //返回数据
}

// 成功
func Sucess(c *gin.Context, data interface{}) {
	if data == nil {
		//数据为空
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(ApiCode.Sucess)
	res.Message = ApiCode.GetMessage(ApiCode.Sucess)
	res.Data = data
	c.JSON(http.StatusOK, res)
}

// 失败
func Failed(c *gin.Context, code int, message string) {
	res := Result{}
	res.Code = code
	res.Message = message
	res.Data = gin.H{}
	c.JSON(http.StatusOK, res)
}
