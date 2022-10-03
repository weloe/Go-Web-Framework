package main

import (
	"fmt"
	"go_web_framework/v1-http-handler/version3/mygin"
	"net/http"
)

func main() {
	myGin := mygin.New()

	myGin.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	myGin.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	})

	myGin.Run(":9999")
}
