package utils

import (
	"e.coding.net/itdesk/weixin/golib/config"
	console_pbs "e.coding.net/itdesk/weixin/golib/pbs/console"
	weixin_pbs "e.coding.net/itdesk/weixin/golib/pbs/weixin"
	"google.golang.org/grpc"
)

var (
	rpcConfig     config.RpcConfig
	weixinClient  weixin_pbs.WeixinServiceClient
	consoleClient console_pbs.UserServiceClient
)

func SetRpcConfig(cfg *config.RpcConfig) {
	rpcConfig = *cfg
}

func NewWeixinClient() error {
	conn, err := grpc.Dial(rpcConfig.WeixinAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	weixinClient = weixin_pbs.NewWeixinServiceClient(conn)
	return nil
}

func NewConsoleClient() error {
	conn, err := grpc.Dial(rpcConfig.ConsoleAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	consoleClient = console_pbs.NewUserServiceClient(conn)
	return nil
}

func WeixinClient() weixin_pbs.WeixinServiceClient {
	if weixinClient == nil {
		NewWeixinClient()
	}
	return weixinClient
}

func ConsoleClient() console_pbs.UserServiceClient {
	if consoleClient == nil {
		NewConsoleClient()
	}
	return consoleClient
}
