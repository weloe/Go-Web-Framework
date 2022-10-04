package mygin

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

//创建路由
func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

//添加路由方法
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

//处理请求
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND：%s\n", c.Path)
	}

}
