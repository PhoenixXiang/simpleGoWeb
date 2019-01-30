package template

import (
	"html/template"
	"io"
)

var t = template.New("app")

func init() {
	t.ParseGlob("./html/*.html")
}

func ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return t.ExecuteTemplate(wr, name, data)
}
