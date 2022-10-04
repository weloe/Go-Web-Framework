package main

import (
	"go_web_framework/v4-route-group-control/mygin"
	"net/http"
)

func main() {
	r := mygin.New()
	r.Get("/index", func(c *mygin.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.Get("/", func(c *mygin.Context) {
			c.HTML(http.StatusOK, "<h1>Hello MyGin</h1>")
		})

		v1.Get("/hello", func(c *mygin.Context) {
			// expect /hello?name=weloe
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.Get("/hello/:name", func(c *mygin.Context) {
			// expect /hello/weloe
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.Post("/login", func(c *mygin.Context) {
			c.JSON(http.StatusOK, mygin.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
