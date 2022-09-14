package config

import "time"

type MysqlConfig struct {
	Addr            string
	Username        string
	Password        string
	DbName          string
	MaxOpenConn     int
	MaxIdleConn     int
	MaxConnLifeTime time.Duration
}

type RedisConfig struct {
	Addr     string
	Password string
}

type WeixinConfig struct {
	WXAppId          string
	WXAppSecret      string
	WXEncodingAESKey string
	WXToken          string
}
