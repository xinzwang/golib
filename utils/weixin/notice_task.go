package weixin

type NoticeTaskProgessData struct {
	First    Keyword `json:"first"`
	Keyword1 Keyword `json:"keyword1"`
	Keyword2 Keyword `json:"keyword2"`
	Keyword3 Keyword `json:"keyword3"`
	Keyword4 Keyword `json:"keyword4"`
	Remark   Keyword `json:"remark"`
}

// 工单处理进度通知
func NoticeTaskProgess(openid string, data *NoticeTaskProgessData, url string) error {
	return TemplateMessage(
		openid,
		"n0P5QhfvrZUEoZbeK5Vnr09e2wX_fwzYOOwKBICUnDM",
		url,
		data,
	)
}

type NoticeTaskTodoData struct {
	First    Keyword `json:"first"`    // 附标题 填问题分类
	Keyword1 Keyword `json:"keyword1"` // 申请单号
	Keyword2 Keyword `json:"keyword2"` // 申请人
	Keyword3 Keyword `json:"keyword3"` // 工单内容
	Keyword4 Keyword `json:"keyword4"` // 申请时间
	Remark   Keyword `json:"remark"`   // 备注
}

// 工单申请待处理通知
func NoticeTaskTodo(openid string, data *NoticeTaskTodoData, url string) error {
	return TemplateMessage(
		openid,
		"	Ba_sGcuOId1LmmBhXPYVtJVOhGkJ-Z_8xXMdHs60ahQ",
		url,
		data,
	)
}
