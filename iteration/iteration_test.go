package iteration

import (
	"fmt"
	"testing"
)

// func TestRepeat(t *testing.T) {
// 	repeated := Repeat("a")
// 	expected := "aaaaa"

// 	if repeated != expected {
// 		t.Errorf("repeated: %s | expected: %s", repeated, expected)
// 	}
// }

// Nos testes de Benchmark, o Go executa a função várias vezes para medir seu desempenho.
// Nesse teste utilizando Builder com WriteString, podemos observar uma melhoria significativa na performance
// Caiu a metade do tempo de execução em comparação com a versão anterior que utilizava concatenação direta de strings.
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 100)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 3)
	fmt.Println(result)
	// Output: aaa
}
