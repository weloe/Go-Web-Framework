package main

import (
	"go_web_framework/v3-dynamic-route/mygin"
	"net/http"
)

func main() {
	r := mygin.New()
	r.Get("/", func(c *mygin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello weloe</h1>")
	})

	r.Get("/hello", func(c *mygin.Context) {
		// expect /hello?name=weloe
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.Get("/hello/:name", func(c *mygin.Context) {
		// expect /hello/weloe
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.Get("/assets/*filepath", func(c *mygin.Context) {
		c.JSON(http.StatusOK, mygin.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
