package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
func EjemploRepetir() {
	repeated := Repeat("d", 7)
	fmt.Println(repeated)
}

func TestRepet(t *testing.T) {
	repeat := Repetir("A", 5)
	expected := "AAAAA"

	if repeat != expected {
		t.Errorf("expected %q but got %q", expected, repeat)
	}
}
