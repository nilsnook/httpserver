package server_test

import (
	"os"
	"server"
	"testing"
)

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}

func TestFileSystemStore(t *testing.T) {
	// database := strings.NewReader(`[
	// 		{"Name": "Cleo", "Wins": 10},
	// 		{"Name": "Chris", "Wins": 33}
	// 	]`)
	database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}
	]`)
	defer cleanDatabase()
	store, err := server.NewFileSystemPlayerStore(database)
	assertNotError(t, err)

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

	t.Run("store wins for existing players", func(t *testing.T) {
		store.RecordWin("Chris")
		want := 34
		got := store.GetPlayerScore("Chris")
		assertScoreEquals(t, want, got)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		store.RecordWin("Pepper")
		want := 1
		got := store.GetPlayerScore("Pepper")
		assertScoreEquals(t, want, got)
	})
}

func assertScoreEquals(t testing.TB, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertNotError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("No error expected, but got one, %v", err)
	}
}
