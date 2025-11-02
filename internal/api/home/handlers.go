package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	response := make(map[string]any)
	response["res"] = fmt.Sprintf("Ola Mundo, %s", name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
