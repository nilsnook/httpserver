package main

import (
	"log"
	"net/http"
	"os"
	"server"
)

const dbFileName = "game.db.json"

func main() {
	// handler := http.HandlerFunc(server.PlayerServer)
	// playerServer := &server.PlayerServer{Store: server.NewInMemoryPlayerStore()}
	// playerServer := server.NewPlayerServer(server.NewInMemoryPlayerStore())
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Problem opening %s %v", dbFileName, err)
	}
	store := &server.FileSystemPlayerStore{Database: db}
	playerServer := server.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", playerServer))
}
