package core

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type MiddlewareHandler func(RequestHandler) RequestHandler

// Log before request be handled
func RequestLogger(inner RequestHandler) RequestHandler {
	return RequestHandler(func(ctx *RequestCtx) (interface{}, error) {
		log.Printf(
			"%s\t%s\t%s",
			ctx.request.RemoteAddr,
			ctx.request.Method,
			ctx.request.RequestURI)

		return inner(ctx)
	})
}

// Log after request be handled
func ResponseLogger(inner RequestHandler) RequestHandler {
	return RequestHandler(func(ctx *RequestCtx) (interface{}, error) {
		start := time.Now()

		result, error := inner(ctx)

		log.Printf(
			"%s\t%s\t%s\t%s",
			ctx.request.RemoteAddr,
			ctx.request.Method,
			ctx.request.RequestURI,
			time.Since(start))

		return result, error
	})
}

func Wrapper(inner RequestHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		inner(&RequestCtx{request, nil, make(map[string]string)})
		fmt.Fprintf(w, "hello")
	})
}
