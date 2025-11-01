package structsMethodsInterfaces

import "testing"

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("Calcula Area", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("Calcular Area do Circulo", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

	// Nesses casos de teste eu posso passar circle e rectangle como Shape
	// Pois eles já se enquadram como um Shape
	// Isso ocorre pois eles receberam a funcao Area no arquivo strucsMethodsInterfaces

	//In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile
}

// Esse modo de teste se chama https://go.dev/wiki/TableDrivenTests
// É muito util quando precisa se testar uma interface ou um struct com varios itens a
// serem validados
func TestArea2(t *testing.T) {
	// Aqui estou criando uma struct anonima
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

//
