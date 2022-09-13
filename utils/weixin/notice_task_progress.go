package weixin

type NoticeTaskProgessData struct {
	First    Keyword `json:"first"`
	Keyword1 Keyword `json:"keyword1"`
	Keyword2 Keyword `json:"keyword2"`
	Keyword3 Keyword `json:"keyword3"`
	Keyword4 Keyword `json:"keyword4"`
	Remark   Keyword `json:"remark"`
}

func NoticeTaskProgess(openid string, data *NoticeTaskProgessData, url string) error {
	return TemplateMessage(
		openid,
		"n0P5QhfvrZUEoZbeK5Vnr09e2wX_fwzYOOwKBICUnDM",
		url,
		data,
	)
}
