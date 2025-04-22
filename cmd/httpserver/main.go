package main

import (
	"log"
	"net/http"

	"github.com/markoc1120/go-application/adapters/httpserver"
	"github.com/markoc1120/go-application/domain/players"
)

const (
	addr = ":8080"
)

func main() {
	server := &httpserver.PlayerServer{Store: players.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(addr, server))
}
