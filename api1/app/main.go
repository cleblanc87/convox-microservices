package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":5000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	profile := Response{"Welcome to api 1"}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
