package main

import (
	"log"
	"net/http"
	"server"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	// handler := http.HandlerFunc(server.PlayerServer)
	server := &server.PlayerServer{Store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
