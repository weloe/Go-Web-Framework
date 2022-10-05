package main

import (
	"fmt"
	"html/template"
	"myGin"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := mygin.New()
	r.Use(mygin.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("./templates/*")
	r.Static("/assets", "./static")
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.Get("/", func(c *mygin.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.Get("/students", func(c *mygin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", mygin.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.Get("/date", func(c *mygin.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", mygin.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
