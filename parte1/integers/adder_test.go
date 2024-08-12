package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	assertSuma(t, sum, expected)

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)

}

func assertSuma(t testing.TB, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("got %q want %q", sum, expected)
	}
}
