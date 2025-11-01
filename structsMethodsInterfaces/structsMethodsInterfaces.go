package structsMethodsInterfaces

import "math"

type Rectangle struct {
	Widht  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rectanlge Rectangle) float64 {
	perimeter := 2 * (rectanlge.Widht + rectanlge.Height)

	return perimeter
}

// func Area(rectangle Rectangle) float64 {
// 	area := rectangle.Widht * rectangle.Height

// 	return area
// }

func (r Rectangle) Area() float64 {
	area := r.Widht * r.Height

	return area
}

// Esse método neste caso nao será utilizado! Usarei o conceito de methodos
// Darei um novo comportamento ao tipo Circle"
// func CircleArea(circle Circle) float64 {
// 	area := circle.Radius * 2

// 	return area
// }"

// (c Circle) é o Receiver
// Ele está recebendo um novo método Area() que calcula a area do Circle e pode ser utilizado como
// Circle.Area()
func (c Circle) Area() float64 {
	area := math.Pi * c.Radius * c.Radius

	return area
}

func (t Triangle) Area() float64 {
	area := (t.Base * t.Height) * 0.5

	return area
}

// A method is a function with a receiver.
// A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

// Methods are very similar to functions but they are called by invoking them on an instance of a particular type.
// Where you can just call functions wherever you like, such as Area(rectangle) you can only call methods on "things".
// A vantagem de usar método é que voce associa ele a uma struct nesse caso, limitando o uso dele a ela
// Um metodo associa um comportamento a um tipo
// por exemplo, Rectangle, e permite chamar r.Area() em vez de Area(r).

// func (c Circle) Area() float64

// (c Circle) := É a struct que recebe a funcao Area
// Area() := É a nova funcao que struct vai receber
// float64 := É o retorno da funcao Area() dentro da struct

// Neste caso nao preciso passar parm pois ja consigo acessar c.Radius com a variavel c
// Porém, caso necessario poderia adicionar um parm ao Area()
//
