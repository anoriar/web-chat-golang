package main

import (
	"chat/internal/adapters"
	"chat/internal/infrastructure/env"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/start", adapters.SocketHandler)
	log.Fatal(http.ListenAndServe(env.Endpoint, nil))
}
