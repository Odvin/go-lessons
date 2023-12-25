package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to person in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to person in Spanish", func(t *testing.T) {
		got := Hello("Lui", "French")
		want := "Bonjour, Lui"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to person in English", func(t *testing.T) {
		got := Hello("Dmytro", "English")
		want := "Hello, Dmytro"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to the world when an empty name is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
