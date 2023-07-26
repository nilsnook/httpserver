package server_test

import (
	"server"
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		store := server.FileSystemPlayerStore{database}

		got := store.GetLeague()
		want := []server.Player{
			{"Cleo", 10},
			{"Chris", 33},
		}
		assertLeague(t, want, got)

		got = store.GetLeague()
		assertLeague(t, want, got)
	})
}
