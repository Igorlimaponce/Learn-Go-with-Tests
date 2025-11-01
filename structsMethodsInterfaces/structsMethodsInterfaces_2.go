package structsMethodsInterfaces

type Shape interface {
	Area() float64
}

// Funcoes definidas dentro de Interface sao "Caracteristicas" necessarias para pertencer aquela interface
// Exemplo se tiver a interface Motorista e dentro tiver 2 funcoes, Dirigir(), Gentil().
// Para uma struct se enquadrar como Motorista ela precisa ter as 2 funcoes definidas para ela

// Como Rectangle e Circle ja possuem um methodo que se chama Area(), em Go ele ja consegue fazer o relacionamento automaticamente
// Ou seja só por ter o Metodo Area, ambos já fazem parte da interface Shape

//
