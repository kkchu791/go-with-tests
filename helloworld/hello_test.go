package main

import "testing"

func TestHello(t *testing.T) { // * is a pointer rather than a copy,  T is a struct that carries all the testing functionality
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Kirk", "")
		want := "Let's begin Kirk"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Let's begin World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Jerome", "french")
		want := "Commencer Jerome"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
