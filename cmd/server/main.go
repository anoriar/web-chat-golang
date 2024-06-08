package main

import (
	"chat/internal/server/adapters"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/start", adapters.SocketHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
