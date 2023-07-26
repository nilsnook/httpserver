package server

import (
	"io"
)

type FileSystemPlayerStore struct {
	Database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	// var league []Player
	// json.NewDecoder(f.Database).Decode(&league)
	league, _ := NewLeague(f.Database)
	return league
}
