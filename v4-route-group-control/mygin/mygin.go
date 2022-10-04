package mygin

import (
	"net/http"
)

// HandlerFunc 定义handler
type HandlerFunc func(c *Context)

type MyGin struct {
	//k 请求方法+路径 v 对应请求处理器
	router *router
}

// New MyGin的构造函数
func New() *MyGin {
	return &MyGin{router: newRouter()}
}

// 添加路由
func (myGin *MyGin) addRoute(method string, pattern string, handler HandlerFunc) {
	myGin.router.addRoute(method, pattern, handler)
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
	c := newContext(w, req)
	myGin.router.handle(c)
}

// Run 启动httpserver
func (myGin *MyGin) Run(addr string) (err error) {
	return http.ListenAndServe(addr, myGin)
}
