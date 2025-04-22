package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/markoc1120/go-application/domain/players"
)

func LeagueSpecification(t testing.TB, store players.PlayerStore) {
	err := store.RecordWin("Lilly")
	assert.NoError(t, err)
	err = store.RecordWin("Lilly")
	assert.NoError(t, err)

	got, _ := store.GetLeague()
	assert.Equal(t, []players.Player{{Name: "Lilly", Wins: 2}}, got)
}
