package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrentMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("want '%q',got '%q'", want, got)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		want := "Hello,girl"
		got := Hello("girl","")

		assertCorrentMessage(t, got, want)
	})

	t.Run("say hello world when a empty name is supplied", func(t *testing.T) {
		want := "Hello,world"
		got := Hello("","")

		assertCorrentMessage(t, got, want)
	})

	t.Run("in spanish",func(t *testing.T){
		want:="Hola,Elodie"
		got:=Hello("Elodie","Spanish")

		assertCorrentMessage(t,got,want)
	})

	t.Run("in french",func(t *testing.T){
		want:="Bonjour,alph"
		got:=Hello("alph","French")

		assertCorrentMessage(t,got,want)
	})
}
