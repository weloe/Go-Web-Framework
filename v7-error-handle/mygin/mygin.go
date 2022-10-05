package mygin

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

// HandlerFunc 定义handler
type HandlerFunc func(c *Context)

type RouterGroup struct {
	prefix      string        // 分组的前缀
	middlewares []HandlerFunc // 中间件
	parent      *RouterGroup  // 父级分组
	myGin       *MyGin        // 所有组共享一个myGin
}

// Group 创建路由分组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	mygin := group.myGin
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		myGin:  mygin,
	}
	mygin.groups = append(mygin.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	// 拼接绝对路径
	absolutePath := path.Join(group.prefix, relativePath)
	// 资源映射
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		// 获取请求的路径参数
		file := c.Param("filepath")
		// 判断该文件是否存在或者是否有权限访问
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// Static 增加静态文件映射
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// 注册handler
	group.Get(urlPattern, handler)
}

type MyGin struct {
	//k 请求方法+路径 v 对应请求处理器
	router *router

	*RouterGroup                     // 嵌套属性，相当于继承了RouterGroup
	groups        []*RouterGroup     // 存储所有路由分组
	htmlTemplates *template.Template //
	funcMap       template.FuncMap   // 自定义模板渲染函数
}

func (myGin *MyGin) SetFuncMap(funcMap template.FuncMap) {
	myGin.funcMap = funcMap
}

func (myGin *MyGin) LoadHTMLGlob(pattern string) {
	myGin.htmlTemplates = template.Must(template.New("").Funcs(myGin.funcMap).ParseGlob(pattern))
}

// New MyGin的构造函数
func New() *MyGin {
	myGin := &MyGin{router: newRouter()}
	myGin.RouterGroup = &RouterGroup{myGin: myGin}
	myGin.groups = []*RouterGroup{myGin.RouterGroup}

	return myGin
}

// Default 默认使用Logger 日志记录 和Recovery 错误处理 中间件
func Default() *MyGin {
	myGin := New()
	myGin.Use(Logger(), Recovery())
	return myGin
}

// 添加路由
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("=== Route registry %4s - %s ===", method, pattern)
	group.myGin.router.addRoute(method, pattern, handler)
}

// Get 添加GET请求路由
func (group *RouterGroup) Get(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// Post 添加POST请求路由
func (group *RouterGroup) Post(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

//实现ServeHTTP方法
func (myGin *MyGin) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc

	//根据前缀找到与该请求匹配的中间件
	for _, group := range myGin.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	c.myGin = myGin
	myGin.router.handle(c)
}

// Run 启动httpserver
func (myGin *MyGin) Run(addr string) (err error) {
	return http.ListenAndServe(addr, myGin)
}
