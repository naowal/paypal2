package app

import "net/http"

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: fixme
	http.NotFound(w, r)
}
