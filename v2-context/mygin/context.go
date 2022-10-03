package mygin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// Context 请求上下文
type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	//请求信息
	Path   string
	Method string
	//响应信息
	StatusCode int
}

// 创建Context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm 获取请求 url 中? 后面的请求参数
// Request.FormValue 方法 可以获取 url 中? 后面的请求参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 获取请求体中的value参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置StatusCode 在header头设置code
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置response header
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String response返回文本
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON response返回json
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)

	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data response返回byte
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML response返回html
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
