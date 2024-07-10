package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{Text: "Hello, JSON!"}
	json.NewEncoder(w).Encode(message)
}

func main() {
	http.HandleFunc("/json", jsonHandler)
	http.ListenAndServe(":8080", nil)
}
