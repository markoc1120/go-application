package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/markoc1120/go-application/adapters"
	"github.com/markoc1120/go-application/adapters/httpserver"
	"github.com/markoc1120/go-application/specifications"
)

func TestHttpServer(t *testing.T) {
	var (
		port    = "8080"
		baseURL = "http://localhost:8080"
		driver  = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)
	if testing.Short() {
		t.Skip()
	}
	adapters.StartDockerServer(t, port, "httpserver")
	specifications.ScoreSpecification(t, driver)
}
