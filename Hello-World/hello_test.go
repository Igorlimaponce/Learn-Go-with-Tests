package main

import (
	"fmt"
	"testing"
)

// *** REQUISITOS DE UM ARQUIVO TESTE ***

// It needs to be in a file with a name like xxx_test.go

// The test function must start with the word Test

// The test function takes one argument only t *testing.T

// To use the *testing.T type, you need to import "testing", like we did with fmt in the other file

// *** ***

// Aqui vai dar um erro
var failNumber = 0

func TestHello(t *testing.T) {
	// Um dos beneficios do subteste é conseguir reaproveitar codigos
	t.Run("Teste com Hello para Pessoas", func(t *testing.T) {
		got := Hello("Igor", "")
		want := "Hello, Igor"

		if got != want {
			t.Errorf("got = %q, want = %q", got, want)
		}
	})

	t.Run("Teste com Hello para Pessoas com valor Vazio", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, Joao"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Em Espanhol", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Em frances", func(t *testing.T) {
		got := Hello("Tom", "French")
		want := "Bonjour, Tom"
		assertCorrectMessage(t, got, want)
	})
}

// Exemplo de reaproveitamento de codigo
func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	// Ele ajuda no debug (O output do terminal retorna a linha em que a funcao que contem este Helper foi chamada (37))
	// Caso eu comente o Helper a linha de retorno será a 47 (Oque é ruim para achar bugs)
	if got != want {
		t.Errorf("got = %q, want = %q", got, want)
	}
}

// Aqui vai passar
// Fiz esse teste primeiro antes de saber que dava para colocar subtestes (Por sinal é muito util)
func TestHello2(t *testing.T) {
	got := Hello("Joao", "")
	want := "Hello, Joao"

	if got != want {
		failNumber += 1
		fmt.Printf("FailNumber: %v\n", failNumber)
		t.Errorf("got = %q, want = %q", got, want)
	}

}
