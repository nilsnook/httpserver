package server

import (
	"io"
)

type FileSystemPlayerStore struct {
	// Database io.Reader
	Database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	// var league []Player
	// json.NewDecoder(f.Database).Decode(&league)
	f.Database.Seek(0, 0)
	league, _ := NewLeague(f.Database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins
}
