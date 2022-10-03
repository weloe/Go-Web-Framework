package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine 处理所有请求的结构体，请求的入口
type Engine struct {
}

// ServeHttp Engine的方法ServeHttp
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}

}

func main() {

	log.Fatal(http.ListenAndServe(":9999", new(Engine)))
}
