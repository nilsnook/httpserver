package server

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	// Database io.Reader
	// Database io.ReadSeeker
	// Database io.Writer
	Database *json.Encoder
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

	// return f.League

	sort.Slice(f.League, func(i, j int) bool {
		return f.League[i].Wins > f.League[j].Wins
	})
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
	// json.NewEncoder(f.Database).Encode(f.League)
	f.Database.Encode(f.League)
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("Problem initializing player DB file, %v", err)
	}
	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("Problem loading player store from file %s, %v", file.Name(), err)
	}
	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&Tape{file}),
		League:   league,
	}, nil
}

func initializePlayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Problem getting file info from file %s, %v", file.Name(), err)
	}
	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}
	return nil
}
