package weixin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"e.coding.net/itdesk/weixin/golib/utils"
)

type templateMsg struct {
	Touser     string      `json:"touser"`
	TemplateId string      `json:"template_id"`
	Url        string      `json:"url"`
	Data       interface{} `json:"data"`
}

type Keyword struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

func TemplateMessage(touser, template_id, url string, msgData interface{}) error {
	s := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", utils.GetAccessToken())
	data := templateMsg{
		Touser:     touser,
		TemplateId: template_id,
		Url:        url,
		Data:       msgData,
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(data)
	_, err := http.Post(s, "application/json; charset=utf-8", buffer)
	return err
}
