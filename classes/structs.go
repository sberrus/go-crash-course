package classes

import "fmt"

// STRUCTS
/*
	Los structs son un tipo de dato que permite representar un dato structurado.
*/

// Mediante el tipado, podemos definir un namespace para definir como va a ser la estructura en cuestion.

type car struct {
	Make  string
	Model string
	Heigh int
	Width int
}

func StructsAnnotations() {
	car := car{
		Make:  "toyota",
		Model: "waooo",
		Heigh: 10,
		Width: 10,
	}

	fmt.Println(car)
}
