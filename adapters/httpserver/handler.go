package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/markoc1120/go-application/domain/players"
)

const (
	playersPath     = "/players"
	leaguePath      = "/league"
	JsonContentType = "application/json"
)

type PlayerServer struct {
	Store players.PlayerStore
	http.Handler
}

func NewPlayerServer(store players.PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.Store = store

	router := http.NewServeMux()
	router.Handle(leaguePath, http.HandlerFunc(p.leagueHandler))
	router.Handle(playersPath+"/", http.HandlerFunc(p.playersHandler))
	p.Handler = router
	return p
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, playersPath+"/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", JsonContentType)
	players, err := p.Store.GetLeague()
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(players)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score, err := p.Store.GetPlayerScore(player)
	if score == 0 || err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	err := p.Store.RecordWin(player)
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
	}
	w.WriteHeader(http.StatusAccepted)
}
