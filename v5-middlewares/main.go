package main

import (
	"go_web_framework/v5-middlewares/mygin"
	"log"
	"net/http"
	"time"
)

func onlyForV2() mygin.HandlerFunc {
	return func(c *mygin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("log error [%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := mygin.New()
	r.Use(mygin.Logger()) // 全局中间件
	r.Get("/", func(c *mygin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello MyGin</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group 中间件
	{
		v2.Get("/hello/:name", func(c *mygin.Context) {
			// expect /hello/MyGin
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
