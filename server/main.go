package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWins(name string) {}

func main() {
	mem := InMemoryPlayerStore{}
	server := &PlayerServer{store: &mem}
	log.Fatal(http.ListenAndServe(":8000", server))
}
