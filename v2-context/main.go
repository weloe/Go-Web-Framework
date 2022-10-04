package main

import (
	"go_web_framework/v2-context/mygin"
	"net/http"
)

func main() {
	r := mygin.New()
	r.Get("/", func(c *mygin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Weloe</h1>")
	})
	r.Get("/hello", func(c *mygin.Context) {
		// expect /hello?name=123
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.Post("/login", func(c *mygin.Context) {
		c.JSON(http.StatusOK, mygin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
