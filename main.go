package main

import (
	"html/template"
	"os"

	"github.com/jockerz/parse-template/parser"
)

type Template struct {
	templates *template.Template
}

func main() {
	tmpl, err := parser.ParseTemplate("templates", "html", nil)
	if err != nil {
		panic(err)
	}

	t := &Template{
		templates: tmpl,
	}

	data := map[string]any{
		"output": "OUTPUT",
	}
	t.templates.ExecuteTemplate(os.Stdout, "lvl1a/lvl2/index.html", data)
}
