package arrays_slices

import (
	"fmt"
)

func SumArray(numbers [5]int) int {

	var sum int

	// Esse for ignora o index com _, util para quando
	// Voce soo precisa dos valores
	for _, number := range numbers {
		sum += number
	}

	// Esse for itera sobre o index
	// ai no caso para somar voce passa a posicao
	// for i := range len(numbers) {
	// 	sum += numbers[i]
	// }

	return sum
}

func SumSlices(numbers []int) int {

	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(n1 []int, n2 []int) []int {

	var result []int
	var sum1 int
	var sum2 int

	for _, i := range n1 {
		sum1 += i
	}

	for _, i := range n2 {
		sum2 += i
	}

	result = append(result, sum1, sum2)

	return result

}

func SumAllAlternative(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = SumSlices(numbers)
	}

	fmt.Printf("numbersToSum = %v", numbersToSum)
	return sums
}

func SumAllAlternative2(numbersToSum ...[]int) []int {
	var sums []int
	// numbers aqui vai ser um slice e vai passar pelo for com cada parte separada do numbersToSum	for _, numbers := range numbersToSum {
	// fmt.Printf("Tipo: %T\n", numbersToSum)
	// fmt.Printf("Quantidade de slices: %d\n", len(numbersToSum))

	for _, numbers := range numbersToSum {
		sums = append(sums, SumSlices(numbers))
	}

	return sums
}

func SumAllTails(slices ...[]int) []int {
	var sums []int
	for _, i := range slices {
		// O objetivo era retornar a soma de todos os elementos menos o primeiro
		// Aqui ele pega a partir do primeiro numero
		// Nesse caso o teste passa normalmente
		i = i[1:]
		sums = append(sums, SumSlices(i))
	}

	return sums
}

// func main() {
// 	SumAllAlternative([]int{1, 2, 3, 4}, []int{8, 7, 6, 5})
// }
