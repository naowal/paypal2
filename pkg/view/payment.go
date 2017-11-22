package view

import "net/http"

type PaymentData struct {
}

type PSuccessData struct {
}

type PDenyData struct {
}

//Payment renders payment view
func Payment(w http.ResponseWriter, r *http.Request, data PaymentData) {
	render(w, r, payment, data)
}

//Success renders pSuccess view
func Success(w http.ResponseWriter, r *http.Request, data PSuccessData) {
	render(w, r, success, data)
}

//Deny renders pDeny view
func Deny(w http.ResponseWriter, r *http.Request, data PDenyData) {
	render(w, r, deny, data)
}
