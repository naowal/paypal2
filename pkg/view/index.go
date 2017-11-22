package view

import (
	"net/http"
)

type IndexData struct {
}

//Index renders index view
func Index(w http.ResponseWriter, r *http.Request, data IndexData) {
	render(w, r, index, data)
}
