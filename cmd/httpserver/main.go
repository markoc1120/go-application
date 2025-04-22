package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/markoc1120/go-application/adapters/httpserver"
	config "github.com/markoc1120/go-application/configuration"
	"github.com/markoc1120/go-application/domain/players"
)

func main() {
	cfg := config.NewConfig()
	server := httpserver.NewPlayerServer(players.NewInMemoryPlayerStore())
	log.Printf("Starting server on :%d", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), server))
}
