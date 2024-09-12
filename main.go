package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerScore struct{}

func (i *InMemoryPlayerScore) GetPlayerScore(name string) int {
	return 123
}
func (i *InMemoryPlayerScore) RecordWin(name string) {}

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
