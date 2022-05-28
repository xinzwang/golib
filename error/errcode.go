package error

const (
	// 用户错误
	FAILED_NO_TOKEN      = 10001 // 无token
	FAILED_INVALID_TOKEN = 10002 // 无效的token

	FAILED_AUTH = 10003 // 验证身份失败
	FAILED_REG  = 10004 // 注册失败

	ERR_NOT_AUTH     = 10100 // 未通过身份验证
	ERR_ALREADY_AUTH = 10101 //已经通过身份验证

	// 应用层错误
	ERR_REQ_DATA     = 20001 // 请求数据结构错误
	ERR_DB_NOT_FOUND = 20002 // 数据库无对应数据

	FAILED_USER_NOT_FOUND = 20003 // 用户尚未录入系统

	// 服务层错误

	// 数据层错误
)
