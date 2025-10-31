package arrays_slices

import (
	"fmt"
	"slices"
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

	t.Run("Teste com 2 Slices de Parametro", func(t *testing.T) {
		// got := SumAll([]int{1, 2}, []int{0, 9})
		// got := SumAllAlternative([]int{1, 2}, []int{0, 9})
		got := SumAllAlternative2([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		// Tomar cuidado ao usar !reflect.DeepEqual
		// Por mais que funcione para alguns casos ela nao bara comparacoes que nao sao do mesmo tipo
		// Tornando perigoso para a aplicacao
		// if !reflect.DeepEqual(got, want) {
		// 	t.Errorf("got %v want %v", got, want)
		// }

		// Esse metodo foi adicionado recentemente e neste caso é melhor por
		// ser focado em slices
		CheckSums(t, got, want)

	})

	t.Run("Teste com a funcao TestSumAllTails, calcula a soma de todos os valores menos do primeiro", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{3, 4, 5})
		want := []int{5, 9}

		CheckSums(t, got, want)
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

func CheckSums(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func CheckSums2(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
