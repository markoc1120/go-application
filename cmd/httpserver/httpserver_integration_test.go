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
	t.Run("score integration test", func(t *testing.T) {
		if testing.Short() {
			t.Skip()
		}
		adapters.StartDockerServer(t, port, "httpserver")
		specifications.ScoreSpecification(t, driver)
	})
	t.Run("league integration test", func(t *testing.T) {
		if testing.Short() {
			t.Skip()
		}
		adapters.StartDockerServer(t, port, "httpserver")
		specifications.LeagueSpecification(t, driver)
	})
}

// func runIntegrationTest(t *testing.T, port string, cmdFolderName string, specFn func(testing.TB, players.PlayerStore)) {
// 	t.Parallel()
// 	if testing.Short() {
// 		t.Skip()
// 	}

// 	baseURL := fmt.Sprintf("http://localhost:%s", port)

// 	driver := httpserver.Driver{
// 		BaseURL: baseURL,
// 		Client: &http.Client{
// 			Timeout: 1 * time.Second,
// 		},
// 	}

// 	adapters.StartDockerServer(t, port, cmdFolderName)
// 	specFn(t, driver)
// }

// func TestHttpServer(t *testing.T) {
// 	t.Run("score integration test", func(t *testing.T) {
// 		runIntegrationTest(t, "8080", "httpserver", specifications.ScoreSpecification)
// 	})

// 	t.Run("league integration test", func(t *testing.T) {
// 		runIntegrationTest(t, "8081", "httpserver", specifications.LeagueSpecification)
// 	})
// }
