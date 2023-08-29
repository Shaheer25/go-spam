package main

import (
    "encoding/json"
    "log"
    "net/http"
)

type Message struct {
    Text string `json:"text"`
}

func main() {
    http.HandleFunc("/api/myendpoint", handleRequest)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    message := Message{Text: "Hello from Go!"}
    json.NewEncoder(w).Encode(message)
}
