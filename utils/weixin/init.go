package weixin

import (
	"fmt"
	url2 "net/url"

	"e.coding.net/itdesk/weixin/golib/config"
)

var (
	cfg config.WeixinConfig
)

func SetWeixinConfig(config *config.WeixinConfig) {
	cfg = *config
}

// 微信网页授权链接
func WeixinAuthUrl(url string) string {
	url = url2.QueryEscape(url)
	url = fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&scope=snsapi_base&response_type=code&redirect_uri=%s", cfg.WXAppId, url)
	return url
}
