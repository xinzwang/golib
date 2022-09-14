package utils

import (
	console_pbs "e.coding.net/itdesk/weixin/golib/pbs/console"
	weixin_pbs "e.coding.net/itdesk/weixin/golib/pbs/weixin"
	"google.golang.org/grpc"
)

var (
	weixinClient  weixin_pbs.WeixinServiceClient
	consoleClient console_pbs.UserServiceClient
)

func NewWeixinClient() error {
	conn, err := grpc.Dial("localhost:8601", grpc.WithInsecure())
	if err != nil {
		return err
	}
	weixinClient = weixin_pbs.NewWeixinServiceClient(conn)
	return nil
}

func NewConsoleClient() error {
	conn, err := grpc.Dial("localhost:8661", grpc.WithInsecure())
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
