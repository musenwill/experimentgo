package main

import (
	"net/http"
	_ "net/http/pprof"
)

func fibRecursive(n int64) int64 {
	if n < 2 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

func main() {
	go func() {
		fibRecursive(10000)
	}()
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
