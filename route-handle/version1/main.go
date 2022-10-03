package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//为路由绑定handler
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

//处理 / 请求
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

//处理
func helloHandler(w http.ResponseWriter, req *http.Request) {
	//遍历请求头
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
}
