package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/markoc1120/go-application/domain/players"
)

func ScoreSpecification(t testing.TB, store players.PlayerStore) {
	got, err := store.GetPlayerScore("Mark")
	assert.NoError(t, err)
	assert.Equal(t, got, 0)

	err = store.RecordWin("Mark")
	assert.NoError(t, err)

	got, err = store.GetPlayerScore("Mark")
	assert.NoError(t, err)
	assert.Equal(t, got, 1)

}
