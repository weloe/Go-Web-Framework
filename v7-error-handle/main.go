package main

import (
	"myGin"
	"net/http"
)

func main() {
	r := mygin.Default()
	r.Get("/", func(c *mygin.Context) {
		c.String(http.StatusOK, "Hello MyGin\n")
	})
	// index out of range for testing Recovery()
	r.Get("/panic", func(c *mygin.Context) {
		names := []string{"MyGin"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
