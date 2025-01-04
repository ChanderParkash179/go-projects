package utils

import (
	"encoding/json"
	"net/http"
)

func APIResponse(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return Encode(w, v)
}

func Encode(w http.ResponseWriter, v any) error {
	return json.NewEncoder(w).Encode(v)
}

func Decode(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}
