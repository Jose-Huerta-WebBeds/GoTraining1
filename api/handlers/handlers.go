package handlers

import (
	"fmt"
	"net/http"

	"github.com/Jose-Huerta-WebBeds/GoTraining1/application/services"
)

// "/"
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d", services.Counter())
}

//Handler returns a handler that can be inserted into an httpServer
func Handler() *http.ServeMux {
	mux := new(http.ServeMux)

	mux.HandleFunc("/count", rootHandler)

	return mux
}
