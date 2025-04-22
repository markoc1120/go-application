package players_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/markoc1120/go-application/adapters/httpserver"
	"github.com/markoc1120/go-application/domain/players"
)

func TestLeague(t *testing.T) {
	wantedLeague := []players.Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}
	store := StubPlayerStore{
		nil,
		nil,
		wantedLeague,
	}
	server := httpserver.NewPlayerServer(&store)
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		assertResponseCode(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response, httpserver.JsonContentType)
	})
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []players.Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}
	return
}

func assertLeague(t testing.TB, got, want []players.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}
