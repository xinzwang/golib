package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Request interface {
	inject(ctx *gin.Context)
	processHttp()
	// process()
}

type BaseApi struct {
	ctx       *gin.Context
	requestId string
	code      int
	msg       string
	errMsg    string
	data      interface{}
}

func (a *BaseApi) inject(ctx *gin.Context) {
	a.ctx = ctx
}

func (a *BaseApi) setErr(code int, msg string, errMsg ...string) *BaseApi {
	a.code = code
	a.msg = msg
	if len(errMsg) > 0 {
		a.errMsg = errMsg[0]
	}
	return a
}

func (a *BaseApi) setCode(code int) *BaseApi {
	a.code = code
	return a
}

func (a *BaseApi) setData(data interface{}) *BaseApi {
	a.data = data
	return a
}

func (a *BaseApi) response() {
	if a.ctx == nil {
		return
	}
	res := gin.H{
		"code": a.code,
		"msg":  a.msg,
	}
	if a.errMsg != "" {
		res["err_msg"] = a.errMsg
	}
	if a.data != nil {
		res["data"] = a.data
	}
	Response(a.ctx, res)
}

func DoHttpProcess(m Request, ctx *gin.Context) { //http处理接口
	//TODO:异常捕获？
	m.inject(ctx)
	m.processHttp()
}

func Response(c *gin.Context, res interface{}) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(res)

	c.Header("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.String(http.StatusOK, buffer.String())
}
