package result

//定义状态码和状态信息

type Codes struct {
	Message map[uint]string
	Sucess  uint
	Failed  uint
}

var ApiCode = &Codes{
	Sucess: 200, //成功的状态码为200
	Failed: 501, //失败的状态吗为501
}

//初始化
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.Sucess: "成功",
		ApiCode.Failed: "失败",
	}
}

//提供获取信息方法
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
