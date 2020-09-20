package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func AssertString(t *testing.T, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("want '%s' but got '%s'", want, got)
	}
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := SpySleeper{Calls: 0}
	Countdown(buffer, &spySleeper)

	got := buffer.String()
	want := `3
2
1
GO!`

	AssertString(t, got, want)

	if spySleeper.Calls != 4 {
		t.Errorf("want call 4 times ,but got %d times", spySleeper.Calls)
	}

}
