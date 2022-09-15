package weixin

// 工单处理进度通知
type NoticeTaskProgessData struct {
	First    Keyword `json:"first"`    // 副标题
	Keyword1 Keyword `json:"keyword1"` // 工单编号
	Keyword2 Keyword `json:"keyword2"` // 工单标题
	Keyword3 Keyword `json:"keyword3"` // 工单状态
	Keyword4 Keyword `json:"keyword4"` // 工单处理人
	Remark   Keyword `json:"remark"`   // 备注
}

func NoticeTaskProgess(openid string, data *NoticeTaskProgessData, url string) error {
	go TemplateMessage(
		openid,
		"n0P5QhfvrZUEoZbeK5Vnr09e2wX_fwzYOOwKBICUnDM",
		url,
		data,
	)
	return nil
}

// 工单申请待处理通知
type NoticeTaskTodoData struct {
	First    Keyword `json:"first"`    // 附标题 填问题分类
	Keyword1 Keyword `json:"keyword1"` // 申请单号
	Keyword2 Keyword `json:"keyword2"` // 申请人
	Keyword3 Keyword `json:"keyword3"` // 工单内容
	Keyword4 Keyword `json:"keyword4"` // 申请时间
	Remark   Keyword `json:"remark"`   // 备注
}

func NoticeTaskTodo(openid string, data *NoticeTaskTodoData, url string) error {
	return TemplateMessage(
		openid,
		"Ba_sGcuOId1LmmBhXPYVtJVOhGkJ-Z_8xXMdHs60ahQ",
		url,
		data,
	)
}

// 工单完成通知
type NoticeTaskFinishData struct {
	First    Keyword `json:"first"`    // 附标题
	Keyword1 Keyword `json:"keyword1"` // 问题描述
	Keyword2 Keyword `json:"keyword2"` // 解决方法
	Keyword3 Keyword `json:"keyword3"` // 反馈时间
	Keyword4 Keyword `json:"keyword4"` // 完成时间
	Remark   Keyword `json:"remark"`   // 备注
}

func NoticeTaskFinish(openid string, data *NoticeTaskProgessData, url string) error {
	return TemplateMessage(
		openid,
		"YmhaIZvb4gOX4tlCJNW3dSReAyaTYDuR6aZUf_u1Z1A",
		url,
		data,
	)
}
