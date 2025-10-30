package tddtypes

func tddtypes() {
	// Em go nós temos alguns tipos de test

	// 1 - test (go test)

	// func TestRepeat(t *testing.T) {
	// 	repeated := Repeat("a")
	// 	expected := "aaaaa"

	// 	if repeated != expected {
	// 		t.Errorf("repeated: %s | expected: %s", repeated, expected)
	// 	}
	// }

	// 2 - Example (go test)

	// func ExampleAdd() {
	// 	sum := Add(1, 6)
	// 	fmt.Println(sum)
	// 	// Output: 6
	// }

	// 3 - Benchmark (go test -bench=)

	// func BenchmarkRepeat(b *testing.B) {
	// 	for b.Loop() {
	// 		Repeat("a")
	// 	}
	// }

	// Quando usamos o Benchamark ela libera a biblioteca Loop
}
