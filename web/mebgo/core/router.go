package core

import (
	"github.com/gorilla/mux"
)

/* Handle request */
type RequestHandler func(ctx *RequestCtx) (interface{}, error)

type Route struct {
	Name           string
	Method         string
	Path           string
	Handler        RequestHandler
	BeforeHandlers []MiddlewareHandler
	AfterHandlers  []MiddlewareHandler
}

var routes = []Route{}

/* Collect custom url route rules */
func RegisterRoute(route Route) {
	bh := []MiddlewareHandler{RequestLogger}
	bh = append(bh, route.BeforeHandlers...)
	route.BeforeHandlers = bh

	ah := route.AfterHandlers
	ah = append(ah, ResponseLogger)
	route.AfterHandlers = ah

	routes = append(routes, route)
}

func LoadRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		requestHandler := route.Handler

		beforeHandlers := reverseMiddlewareHandlers(route.BeforeHandlers)
		for _, h := range beforeHandlers {
			requestHandler = h(requestHandler)
		}

		for _, h := range route.AfterHandlers {
			requestHandler = h(requestHandler)
		}

		httpHandler := Wrapper(requestHandler)
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(httpHandler)
	}
	return router
}

func reverseMiddlewareHandlers(handlers []MiddlewareHandler) []MiddlewareHandler {
	last := len(handlers) - 1
	for i := 0; i < len(handlers)/2; i++ {
		handlers[i], handlers[last-i] = handlers[last-i], handlers[i]
	}
	return handlers
}
