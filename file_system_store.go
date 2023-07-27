package server

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	// Database io.Reader
	// Database io.ReadSeeker
	Database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {
	// var league []Player
	// json.NewDecoder(f.Database).Decode(&league)
	f.Database.Seek(0, 0)
	league, _ := NewLeague(f.Database)
	return league
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
	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	// for i, player := range league {
	// 	if player.Name == name {
	// 		league[i].Wins++
	// 	}
	// }
	player := league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}
	f.Database.Seek(0, 0)
	json.NewEncoder(f.Database).Encode(league)
}
