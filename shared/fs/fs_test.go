package fs

import (
	"testing"
	"testing/fstest"
)

// https://stackoverflow.com/questions/43912124/example-code-for-testing-the-filesystem-in-golang
func TestExample(t *testing.T) {
	m := fstest.MapFS{
		"hello.txt": {
			Data: []byte("hello, world"),
		},
	}
	b, e := m.ReadFile("hello.txt")
	if e != nil {
		panic(e)
	}
	println(string(b) == "hello, world")
}
