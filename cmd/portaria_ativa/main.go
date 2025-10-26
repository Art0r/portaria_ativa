package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Art0r/portaria_ativa/internal"
	web "github.com/Art0r/portaria_ativa/internal/web/home"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	web.SetHomeRoutes(r)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(internal.AppConfig.StaticDir))))

	server := &http.Server{
		Handler:           r,
		Addr:              "127.0.0.1:8000",
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}

	fmt.Println("Server running on http://127.0.0.1:8000")
	log.Fatal(server.ListenAndServe())
}
