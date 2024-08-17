package handler

import (
	"fmt"
	"net/http"
)

var contador int

func Handler(w http.ResponseWriter, r *http.Request) {
	contador++
	fmt.Fprintf(w, "<p>%v: %v</p>", contador, r.URL)
}
