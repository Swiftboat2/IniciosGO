package main

import (
	"reflect"
	"testing"
)

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

		got := Suma(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})
}

func TestSumaTodo(t *testing.T) {
	got := SumaTodo([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}
func TestSumaTodoFinales(t *testing.T) {
	t.Run("suma el final del array", func(t *testing.T) {
		got := SumaTodoFinal([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("suma slices vacias", func(t *testing.T) {
		got := SumaTodoFinal([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
