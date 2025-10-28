package main

import (
	"testing"
)

// *** REQUISITOS DE UM ARQUIVO TESTE ***

// It needs to be in a file with a name like xxx_test.go

// The test function must start with the word Test

// The test function takes one argument only t *testing.T

// To use the *testing.T type, you need to import "testing", like we did with fmt in the other file

// *** ***

// Aqui vai dar um erro
func TestHello(t *testing.T) {
	got := Hello("Igor")
	want := "Hello, World!!"

	if got != want {
		t.Errorf("got = %q, want = %q", got, want)
	}
}

// Aqui vai passar
func TestHello2(t *testing.T) {
	got := Hello("Joao")
	want := "Hello World! Joao"

	if got != want {
		t.Errorf("got = %q, want = %q", got, want)
	}
}
