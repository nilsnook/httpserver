package server

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	// Database io.Reader
	// Database io.ReadSeeker
	Database io.Writer
	League   League
}

func (f *FileSystemPlayerStore) GetLeague() League {
	// var league []Player
	// json.NewDecoder(f.Database).Decode(&league)

	// The problem with this approach is that every time
	// we call GetLeague, we are reading the entire file
	// and parsing it into json

	// f.Database.Seek(0, 0)
	// league, _ := NewLeague(f.Database)
	// return league

	return f.League
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	// var wins int
	// for _, player := range f.GetLeague() {
	// 	if player.Name == name {
	// 		wins = player.Wins
	// 		break
	// 	}
	// }
	// return wins

	// player := f.GetLeague().Find(name)
	player := f.League.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	// league := f.GetLeague()

	// for i, player := range league {
	// 	if player.Name == name {
	// 		league[i].Wins++
	// 	}
	// }

	player := f.League.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.League = append(f.League, Player{name, 1})
	}
	// f.Database.Seek(0, 0)
	json.NewEncoder(f.Database).Encode(f.League)
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, io.SeekStart)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		Database: &Tape{database},
		League:   league,
	}
}
