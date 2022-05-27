package error

const (
	// 用户错误
	NOT_LOGIN   = 10001 // 未登录
	FAILED_AUTH = 10002 // 验证身份失败
	FAILED_REG  = 10003 // 注册失败

	// 应用层错误
	ERR_REQ_DATA     = 20001 // 请求数据结构错误
	ERR_DB_NOT_FOUND = 20002 // 数据库无对应数据

	// 服务层错误

	// 数据层错误
)
