package app

import "net/http"

//Mount mounts handlers to mux
func Mount(mux *http.ServeMux) {

	mux.Handle("/", makeGetHandler(indexHandler))
	mux.Handle("/payment", makeGetHandler(PaymentHandler))
	mux.Handle("/success", makeGetHandler(PSuccessHandler))
	mux.Handle("/deny", makeGetHandler(PDenyHandler))

}

func makeGetHandler(get http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet, http.MethodHead:
			get.ServeHTTP(w, r)
		default:
			notFoundHandler(w, r)
		}
	})
}
