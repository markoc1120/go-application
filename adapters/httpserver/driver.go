package httpserver

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Driver struct {
	BaseURL string
	Client  *http.Client
}

func (d Driver) GetPlayerScore(name string) (int, error) {
	res, err := d.Client.Get(d.BaseURL + fmt.Sprintf("%s/%s", playersPath, name))
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	scoreBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	score, err := strconv.Atoi(string(scoreBytes))
	if err != nil {
		return 0, err
	}
	return score, nil
}

func (d Driver) RecordWin(name string) error {
	res, err := d.Client.Post(d.BaseURL+fmt.Sprintf("%s/%s", playersPath, name), "application/json", nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
