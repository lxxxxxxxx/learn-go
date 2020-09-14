package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello girl"
	got := Hello("girl")

	if got != want {
		t.Errorf("\n got '%q' \n want '%q'", got, want)
	}
}
