package api

import (
	"github.com/Art0r/portaria_ativa/internal/api/middlewares"
	"github.com/gorilla/mux"
)

func SetHomeRoutes(r *mux.Router) {
	homeRoutes := r.PathPrefix("/home").Subrouter()
	homeRoutes.HandleFunc("", HomeHandler)
	homeRoutes.Use(middlewares.DefaultMiddleware)
}
