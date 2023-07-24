package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.showScore(w, r)
	case http.MethodPost:
		p.processWin(w, r)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.Store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	p.Store.RecordWin("Bob")
	w.WriteHeader(http.StatusAccepted)
}

// func GetPlayerScore(name string) string {
// 	switch name {
// 	case "Pepper":
// 		return "20"
// 	case "Floyd":
// 		return "10"
// 	default:
// 		return "No such player"
// 	}
// }
