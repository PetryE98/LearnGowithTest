package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {

			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Eric", "English")
		want := "Hello, Eric"

		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Eric", "Spanish")
		want := "Hola,Eric"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Eric", "French")
		want := "salut,Eric"
		assertCorrectMessage(t, got, want)

	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {

		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}
