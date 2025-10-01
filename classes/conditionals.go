package classes

import "fmt"

func ConditionalNotes() {
	// CONDICIONALES

	value := 25000

	if value < 10 {
		fmt.Printf("The value is less than 10. Value: %d\n", value)
	} else if value >= 10 && value <= 15 {
		fmt.Printf("The number is between 10..15. Value: %d\n", value)
	} else {
		fmt.Printf("The value is bigger than 15. Value %d\n", value)
	}

	// Condicionales recortada
	/*
		En este caso lo que debemos tomar en cuenta
		es que al momento de asignar un valor a una
		variable, podemos usar esta sintaxis recortada.

		Seguido del ";" podemos definir la condicion el cual quedaría de la siguiente manera:

		if <Asignación variable>; <condicion> {...}
	*/

	// lenght solo es accesible dentro de este bloque
	// añade más seguridad
	if lenght := 10; lenght < 15 {
		fmt.Println("Wow, i just did this")
	} else {
		fmt.Println("Or not :v")
	}
}
