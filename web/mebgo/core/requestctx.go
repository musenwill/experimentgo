package core

import "net/http"

type RequestCtx struct {
	request *http.Request
	param   interface{}
	session map[string]string
}

func (ctx *RequestCtx) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *RequestCtx) SetParam(param interface{}) {
	ctx.param = param
}

func (ctx *RequestCtx) SessionGet(key string) (string, bool) {
	value, exists := ctx.session[key]
	return value, exists
}

func (ctx *RequestCtx) SessionSet(key string, value string) {
	ctx.session[key] = value
}
