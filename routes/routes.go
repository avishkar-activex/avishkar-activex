package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Register(r *mux.Router) {
	r.HandleFunc("/healthcheck", healthCheck).Methods(http.MethodGet)
	r.HandleFunc("/v1/auth", authenticate).Methods(http.MethodPost)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

}

func authenticate(w http.ResponseWriter, r *http.Request) {

}
