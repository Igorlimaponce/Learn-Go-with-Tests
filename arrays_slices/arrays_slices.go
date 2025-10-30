package arrays_slices

func SumArray(numbers [5]int) int {

	var sum int

	// Esse for ignora o index com _, util para quando
	// Voce soo precisa dos valores
	for _, number := range numbers {
		sum += number
	}

	// Esse for itera sobre o index
	// ai no caso para somar voce passa a posicao
	for i := range len(numbers) {
		sum += numbers[i]
	}

	return sum
}

func SumSlices(numbers []int) int {

	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}
