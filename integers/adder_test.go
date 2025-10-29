package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Teste que deveria dar erro", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 5

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("Teste que deveria passar", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 6)
	fmt.Println(sum)
	// Output: 6
}
