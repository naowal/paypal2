package app

import (
	"net/http"

	"github.com/naowal/paypal2/pkg/view"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	view.Payment(w, r, view.PaymentData{})
}

func PSuccessHandler(w http.ResponseWriter, r *http.Request) {
	view.Success(w, r, view.PSuccessData{})
}

func PDenyHandler(w http.ResponseWriter, r *http.Request) {
	view.Deny(w, r, view.PDenyData{})
}
