package config

// 本服务 运行时环境配置
type RuntimeConfig struct {
	Domain   string
	HttpPort string
	RpcPort  string
}

// 本服务调用其它服务的rpc的ip地址
type RpcConfig struct {
	WeixinAddr  string
	ConsoleAddr string
}
