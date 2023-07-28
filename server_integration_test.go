package server_test

import (
	"net/http"
	"net/http/httptest"
	"server"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	// store := server.NewInMemoryPlayerStore()

	// Earlier we were not getting errors before because
	// we always just ignored them -> line 19
	// database, cleanDatabase := createTempFile(t, "")

	database, cleanDatabase := createTempFile(t, `[]`)
	defer cleanDatabase()
	store, err := server.NewFileSystemPlayerStore(database)
	assertNotError(t, err)
	playerServer := server.NewPlayerServer(store)
	player := "Pepper"

	playerServer.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	playerServer.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	playerServer.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		want := []server.Player{
			{"Pepper", 3},
		}
		got := getLeagueFromResponse(t, response.Body)
		assertLeague(t, want, got)
	})
}
