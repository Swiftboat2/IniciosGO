package main

import "testing"

func TestSum(t *testing.T) {
t.Run("coleccion de 5 numeros", func(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Suma(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
})


t.Run("colleccion de cualquier tamaño", func(t *testing.T) {
	numbers := []int{1, 2, 3}

	got:= Suma(numbers)
	want := 6

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}


})


}
