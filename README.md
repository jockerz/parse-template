<p class="center">
  <a href="https://codecov.io/gh/jockerz/parse-template" target="_blank">
    <img src="https://img.shields.io/codecov/c/github/jockerz/parse-template?color=%2334D058" alt="Coverage">
  </a>
  <a href="https://goreportcard.com/report/github.com/jockerz/parse-template" target="_blank">
    <img src="https://goreportcard.com/badge/github.com/jockerz/parse-template" alt="Coverage">
  </a>
   <a href="https://github.com/jockerz/parse-template/actions">
    <img src='https://github.com/jockerz/parse-template/actions/workflows/go.yml/badge.svg' alt='Test Status' />
  </a>
</p>


# Introduction

Namespaced `template.Template` files by directory made easier.

# Usage

```
go get "github.com/jockerz/parse-template/parser"
```


## Example

Template Contents
```
$ tree templates/
templates/
├── lvl1a
│   ├── index.html
│   ├── lvl2
│   │   └── index.html
│   └── lvl2_empty
└── parts
    └── include_me.html
```

Content of `templates/lvl1a/index.html`
```
Content of `lvl1a/index.html`
```

Content of `templates/lvl1a/lvl2/index.html`
```
Content of `lvl1a/lvl2/index.html`

Read data from execute {{ .output }}

{{ template "parts/include_me.html" . }}
```

Content of `templates/parts/include_me.html`
```
TEXT from "parts/include_me.html"
```

`main.go`
```go
package main

import (
	"fmt"
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
	err = t.templates.ExecuteTemplate(os.Stdout, "lvl1a/lvl2/index.html", data)
	fmt.Println(err.Error())
}
```

Result of `go run main.go`
```shell
$ go run main.go
Content of `lvl1a/lvl2/index.html`

Read data from execute OUTPUT

TEXT from "parts/include_me.html"

```

## Test

```
go test -v ./...
```
