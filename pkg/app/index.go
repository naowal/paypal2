package app

import (
	"net/http"

	"github.com/naowal/paypal2/pkg/view"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	view.Index(w, r, view.IndexData{})
}
