package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("diciendo hola a la gente", func(t *testing.T) {

		got := Hello("Chris")
		want := "Hola Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("deci 'Hola mundo' cuando se entregue un string vacio", func(t *testing.T) {
		got := Hello("")
		want := "Hola Mundo"

		if got != want {
			t.Errorf("got %q want %q", got, want)

		}
	})

}
