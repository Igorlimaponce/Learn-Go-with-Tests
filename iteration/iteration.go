package iteration

import "strings"

const repeatCount = 5

// Ao inves de concatenar strings diretamente (x += str) que é ineficiente,
// uma abordagem melhor seria usar strings.Builder ou strings.Repeat.
// Fazendo desta forma temos uma performance melhor.
// https://pkg.go.dev/strings#Builder
func Repeat(str string, repeatTimes int) string {
	var repeated strings.Builder
	for i := 0; i < repeatTimes; i++ {
		repeated.WriteString(str)
	}

	return repeated.String()
}
