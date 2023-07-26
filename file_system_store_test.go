package server_test

import (
	"server"
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
	store := server.FileSystemPlayerStore{database}

	t.Run("league from a reader", func(t *testing.T) {
		want := []server.Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		got := store.GetLeague()
		assertLeague(t, want, got)

		got = store.GetLeague()
		assertLeague(t, want, got)
	})

	t.Run("get player score", func(t *testing.T) {
		want := 33
		got := store.GetPlayerScore("Chris")
		assertScoreEquals(t, want, got)
	})
}

func assertScoreEquals(t testing.TB, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
