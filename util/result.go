package util

import (
	"RuoYi-Go/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 这里得 `json`是用于json得反序列化使用得
// 并且 Data 字段 使用得是interface{} 是用于存储任意类型的数据
type Result struct {
	Time time.Time   `json:"time"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 成功返回消息
func Success(c *gin.Context, data interface{}) {
	// 如果传入得参数为空 则默认返回空对象 gin框架提供快捷方式 gin.h{}
	if data == nil {
		data = gin.H{}
	}

	//初始化结构体 结构体得字段都将被赋值为0
	res := Result{}

	// 对字段赋值
	res.Time = time.Now()
	res.Code = int(types.ApiCode.SUCCESS)
	res.Msg = types.ApiCode.GetMessage(types.ApiCode.SUCCESS)
	res.Data = data
	// 生成并相应json数据
	c.JSON(http.StatusOK, res)
}

// 返回错误信息
func Error(c *gin.Context, code int, message string) {
	res := Result{}
	res.Time = time.Now()
	res.Code = code
	res.Msg = message
	res.Data = gin.H{}
	c.JSON(http.StatusInternalServerError, res)
}
