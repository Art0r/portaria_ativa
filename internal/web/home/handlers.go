package web

import (
	"html/template"
	"math/rand"
	"net/http"
)

type HomeData struct {
	Msg string
}

type CountData struct {
	Value uint32
}

var templates = template.Must(template.ParseGlob("./public/views/home/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	err := templates.ExecuteTemplate(w, "index.html", HomeData{
		Msg: "Ola Mundoi",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	err := templates.ExecuteTemplate(w, "count.html", CountData{
		Value: rand.Uint32(),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
