package config

type WeixinConfig struct {
	WXAppId          string
	WXAppSecret      string
	WXEncodingAESKey string
	WXToken          string
}

type AlibabaConfig struct {
	AccessKeyId     string
	AccessKeySecret string
}
