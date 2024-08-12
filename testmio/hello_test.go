package main

import "testing"

func TestMain(t *testing.T) {
	t.Run("saludando en ingles", func(t *testing.T) {

		got := Yo("", "")
		want := "Hello Lauty Pelozo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saludando en español", func(t *testing.T) {

		got := Yo("", Español)
		want := "Hola Lauty Pelozo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saludando frances", func(t *testing.T) {
		got := Yo("", Bounjour)
		want := "Bounjour Lauty Pelozo"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
