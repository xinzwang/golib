package utils

import (
	"context"
	"log"
	"sync"
	"time"

	weixin_pbs "e.coding.net/itdesk/weixin/golib/pbs/weixin"

	"google.golang.org/grpc"
)

var (
	AccessToken     string
	accessExpiresIn int64
	JsapiTicket     string
	jsapiExpiresIn  int64

	lock sync.Mutex
)

func init() {
	initTickets()
}

func initTickets() {
	conn, err := grpc.Dial("localhost:8601", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := weixin_pbs.NewWeixinServiceClient(conn)
	res, err := client.TicketRpc(context.Background(), &weixin_pbs.TicketReq{
		TicketType: weixin_pbs.TicketType_access_token,
	})
	AccessToken = res.GetTicket()
	accessExpiresIn = res.GetExpiresIn()

	res, err = client.TicketRpc(context.Background(), &weixin_pbs.TicketReq{
		TicketType: weixin_pbs.TicketType_jsapi_ticket,
	})
	JsapiTicket = res.GetTicket()
	jsapiExpiresIn = res.GetExpiresIn()
}

func GetAccessToken() string {
	if accessExpiresIn < time.Now().Unix() {
		lock.Lock()
		defer lock.Unlock()
		if accessExpiresIn < time.Now().Unix() {
			initTickets()
		}
	}
	return AccessToken
}

func GetJsApiTIcket() string {
	if jsapiExpiresIn < time.Now().Unix() {
		lock.Lock()
		defer lock.Unlock()
		if jsapiExpiresIn < time.Now().Unix() {
			initTickets()
		}
	}
	return JsapiTicket
}
