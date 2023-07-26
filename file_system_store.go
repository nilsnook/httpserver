package server

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	Database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	var league []Player
	json.NewDecoder(f.Database).Decode(&league)
	return league
}
