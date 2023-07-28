package server

import "io"

type Tape struct {
	file io.ReadWriteSeeker
}

func (t *Tape) Write(p []byte) (n int, err error) {
	t.file.Seek(0, io.SeekStart)
	return t.file.Write(p)
}
