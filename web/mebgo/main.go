package main

import (
	"encoding/json"
	"fmt"

	"github.com/musenwill/experimentgo/web/mebgo/core"
)

func main() {
	core.RegisterRoute(core.Route{Name: "index", Method: "GET", Path: "/", Handler: index})
	core.RegisterRoute(core.Route{Name: "index", Method: "POST", Path: "/", Handler: post})
	core.Run()
}

func index(ctx *core.RequestCtx) (interface{}, error) {
	request := ctx.GetRequest()
	request.ParseForm()
	fmt.Println(request.Form)
	return nil, nil
}

func post(ctx *core.RequestCtx) (interface{}, error) {
	request := ctx.GetRequest()
	param := make(map[string]string)
	json.NewDecoder(request.Body).Decode(&param)
	fmt.Println(param)
	return nil, nil
}
