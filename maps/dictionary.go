package maps

import "errors"

type Dictionary map[string]string

// Map: Tem um campo key e um campo value
// Cada key é unica e faz referencia a um determinado valor
// func Search(dictionary map[string]string, name string) string {
// 	return dictionary[name]
// }

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	// Aqui utilizamos uma propriedade interessante do map, ele consegue retornar dois valores
	// o promeiro é o valor em si daquela chave, e o segundo um boolean falando se a chave foi encontrada ou nao
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
