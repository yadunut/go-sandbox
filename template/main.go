package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/gobuffalo/packr"
)

var (
	templates *template.Template
)

func init() {
	box := packr.NewBox("./template")
	templates = template.New("")
	box.Walk(func(s string, _ packr.File) error {
		template.Must(templates.New(s).Parse(box.String(s)))
		return nil
	})
}

func main() {
	for i, v := range templates.Templates() {
		fmt.Println(i, v.Execute(os.Stdout, nil))
	}
}
