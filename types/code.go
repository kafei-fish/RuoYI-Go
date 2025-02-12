package types

type Codes struct {
	SUCCESS          uint
	FAILED           uint
	GENERATETOKEN    uint
	NOAUTH           uint
	AUTHFAILED       uint
	AUTHFORMATERROR  uint
	INVALIDTOKEN     uint
	NOSUCHID         uint
	CREATEUSERFAILED uint
	LCAKPARAMETERS   uint
	CONVERTFAILED    uint
	NOSUCHNAME       uint
	EXISTSNAME       uint
	Message          map[uint]string
}

// 初始化APi code

var ApiCode = &Codes{
	SUCCESS:          200,
	FAILED:           0,
	AUTHFAILED:       4001,
	GENERATETOKEN:    4002,
	NOAUTH:           4003,
	AUTHFORMATERROR:  4004,
	INVALIDTOKEN:     4005,
	NOSUCHID:         1001,
	CREATEUSERFAILED: 2001,
	LCAKPARAMETERS:   3001,
	CONVERTFAILED:    3002,
	NOSUCHNAME:       5001,
	EXISTSNAME:       5002,
}

// 初始化 错误码
// 为啥不用make 因为在初始化map已经赋值参数了
// make 函数通常用于创建和初始化 map、slice 和 channel 使用 make 函数并指定初始容量
func init() {
	ApiCode.Message = map[uint]string{
		ApiCode.SUCCESS:          "成功",
		ApiCode.FAILED:           "失败",
		ApiCode.GENERATETOKEN:    "生成Token失败",
		ApiCode.AUTHFAILED:       "鉴权失败",
		ApiCode.NOAUTH:           "请求头中auth为空",
		ApiCode.AUTHFORMATERROR:  "请求头中auth格式有误",
		ApiCode.INVALIDTOKEN:     "无效的Token",
		ApiCode.NOSUCHID:         "id不存在",
		ApiCode.CREATEUSERFAILED: "用户创建失败",
		ApiCode.LCAKPARAMETERS:   "缺少参数",
		ApiCode.CONVERTFAILED:    "参数类型转换报错",
		ApiCode.NOSUCHNAME:       "根据名称查不到数据",
		ApiCode.EXISTSNAME:       "名称重复",
	}
}

/*
*
获取消息字符串 接收者是一个Codes类型的指针。这个方法只能被Codes类型的指针对象调用：
如果接收者是指针类型，则该方法只能被指针对象调用。
*/
func (c *Codes) GetMessage(code uint) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
