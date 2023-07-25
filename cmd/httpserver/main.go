package main

import (
	"log"
	"net/http"
	"server"
)

func main() {
	// handler := http.HandlerFunc(server.PlayerServer)
	// playerServer := &server.PlayerServer{Store: server.NewInMemoryPlayerStore()}
	playerServer := server.NewPlayerServer(server.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", playerServer))
}
