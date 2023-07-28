package server_test

import (
	"io"
	"os"
	"server"
	"testing"
)

func TestTapeWrite(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &server.Tape{file.(*os.File)}
	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	want := "abc"
	got := string(newFileContents)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
