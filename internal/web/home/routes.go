package web

import "github.com/gorilla/mux"

func SetHomeRoutes(r *mux.Router) {
	homeRoutes := r.PathPrefix("/home").Subrouter()
	homeRoutes.HandleFunc("", HomeHandler)
	homeRoutes.HandleFunc("/counter", CountHandler)
}
