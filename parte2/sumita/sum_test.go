package sumita

import (
	"testing"
)

func TestSuma(t *testing.T) {
	numeros := [5]int{1, 2, 3, 4, 5}

	got := Suma(numeros)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numeros)
	}
}
