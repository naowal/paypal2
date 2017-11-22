package view

import (
	"bytes"
	"net/http"

	"github.com/acoshift/session"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

var (
	m = minify.New()
)

//init initialize minify function
func init() {
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/javascript", js.Minify)
}

// render renders given template
func render(w http.ResponseWriter, r *http.Request, tpl int, data interface{}) {
	t, ok := tmpl[tpl]
	if !ok {
		// developer error
		panic("view: template not found")
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")
	buf := bytes.Buffer{}
	err := t.Execute(&buf, data)
	if err != nil {
		panic("view: template execute error")
	}
	m.Minify("text/html", w, &buf)
	session.Get(r.Context(), "sess").Flash().Clear()
}
