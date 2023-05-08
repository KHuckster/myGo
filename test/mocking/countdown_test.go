package mocking

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	sleeper := &SpySleeper{}
	Countdown(&buffer, sleeper)
	got := buffer.String()
	want := `3
2
1
GO!`
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if sleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", sleeper.Calls)
	}

}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i >= 1; i-- {
		sleeper.Sleep()
		_, _ = fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	_, _ = fmt.Fprint(out, "GO!")
}
