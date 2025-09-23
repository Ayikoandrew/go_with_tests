package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWins(string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if r.Method == http.MethodPost {
		p.processWin(w)
		return
	}
	if r.Method == http.MethodGet {
		p.showWin(w, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
	p.store.RecordWins("Bob")
	w.WriteHeader(http.StatusAccepted)
}
func (p *PlayerServer) showWin(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}
