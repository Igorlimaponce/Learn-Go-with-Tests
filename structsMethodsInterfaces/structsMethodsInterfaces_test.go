package structsMethodsInterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Primeiro teste com Perimeter", func(t *testing.T) {
		// Aqui estou usando a struct criada,
		// Parms Widht e Height
		rectangle := Rectangle{10.5, 12.3}
		got := Perimeter(rectangle)
		want := 45.6

		if got != want {
			t.Errorf("got: %.2f   want: %.2f", got, want)
		}
	})

	t.Run("Calcula Area", func(t *testing.T) {
		rectangle := Rectangle{10.3, 10.2}
		got := rectangle.Area()
		want := 105.06

		if got != want {
			t.Errorf("got: %.2f want: %.2f", got, want)
		}
	})

	t.Run("Calcular Area do Circulo", func(t *testing.T) {
		circle := Circle{10.3}
		got := circle.Area()
		want := 333.2915646193412

		if got != want {
			t.Errorf("got: %g, want: %g", got, want)
		}
		// Nesse caso usei o %g para formatacao pois o numero tem varias casas decimais
		// Sendo assim o %g mostra os numeros de forma mais precisa
	})
}

//
