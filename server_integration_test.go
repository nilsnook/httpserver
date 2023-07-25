package server_test

import (
	"net/http"
	"net/http/httptest"
	"server"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := server.NewInMemoryPlayerStore()
	// server := server.PlayerServer{store}
	server := server.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
