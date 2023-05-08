package di

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestDi(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "zk")

	got := buffer.String()
	want := "Hello, zk"
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func Greet(write io.Writer, name string) {
	_, _ = fmt.Fprintf(write, "Hello, %s", name)
}
