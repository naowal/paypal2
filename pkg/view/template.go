package view

import (
	"html/template"
	"path/filepath"
)

const (
	_ = iota
	index
	payment
	success
	deny
)

var tmpl = map[int]*template.Template{
	index:   load("index", "layout/root"),
	payment: load("payment", "layout/root"),
	success: load("success", "layout/root"),
	deny:    load("deny", "layout/root"),
}

const templateDir = "template"

// joinTemplateDir joins all given filenames with template dir and extension ".tmpl"
func joinTemplateDir(filenames ...string) []string {
	rs := make([]string, len(filenames))
	for i, x := range filenames {
		rs[i] = filepath.Join(templateDir, x) + ".tmpl"
	}
	return rs
}

// load loads template or panic
func load(filenames ...string) *template.Template {
	return template.Must(template.ParseFiles(joinTemplateDir(filenames...)...)).Lookup("root")
}
