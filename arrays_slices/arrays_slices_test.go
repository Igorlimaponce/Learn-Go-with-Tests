package arrays_slices

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Teste com Array (Vetor)", func(t *testing.T) {
		// Array (Vetor)
		numbers := [5]int{1, 2, 3, 4, 5}

		got := SumArray(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("Teste com Slices", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := SumSlices(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func ExampleSumArray() {
	numbers := [5]int{1, 2, 3, 4, 5}
	result := SumArray(numbers)
	fmt.Println(result)
	// Output: 15
}

func ExampleSumSlices() {
	numbers := []int{1, 2, 3, 4, 5}
	result := SumSlices(numbers)
	fmt.Println(result)
	// Output: 15
}
