package fileWrite_test.go

import "testing"

func TestFileWrite(t *testing.T) {
    result := Hello("World!")
    want := "Hi, Arei. Welcome!"
    if result != want {
        t.Errorf("fileWrite.Hello() = %q want %q", result, want)
    }
}