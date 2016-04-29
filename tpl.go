package main

import (
	"io"
	"text/template"

	"github.com/Masterminds/sprig"
)

func createTpl(tplName, fileName string) (*template.Template, error) {
	return template.New(tplName).Funcs(sprig.TxtFuncMap()).ParseFiles(fileName)
}

func renderTpl(tpl *template.Template, w io.Writer, data interface{}) error {
	return tpl.Execute(w, data)
}
