package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("diciendo hola a la gente", func(t *testing.T) {

		got := Hello("Chris")
		want := "Hola Chris"
		assertCorrectMessage(t, got, want)

	})

	t.Run("deci 'Hola mundo' cuando se entregue un string vacio", func(t *testing.T) {
		got := Hello("")
		want := "Hola Mundo"
		assertCorrectMessage(t, got, want)
	})
}

//en got, want string se puede poner una sola vez string al final
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
