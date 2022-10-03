package mygin

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义handler
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type MyGin struct {
	//k 请求方法+路径 v 对应请求处理器
	router map[string]HandlerFunc
}

// New MyGin的构造函数
func New() *MyGin {
	return &MyGin{router: make(map[string]HandlerFunc)}
}

// 添加路由
func (myGin *MyGin) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	myGin.router[key] = handler
}

// Get 添加GET请求路由
func (myGin *MyGin) Get(pattern string, handler HandlerFunc) {
	myGin.addRoute("GET", pattern, handler)
}

// Post 添加POST请求路由
func (myGin *MyGin) Post(pattern string, handler HandlerFunc) {
	myGin.addRoute("POST", pattern, handler)
}

//实现ServeHTTP方法
func (myGin *MyGin) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	//遍历myGin的map
	if handler, ok := myGin.router[key]; ok {
		//如果有对应的value
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}

}

// Run 启动httpserver
func (myGin *MyGin) Run(addr string) (err error) {
	return http.ListenAndServe(addr, myGin)
}
