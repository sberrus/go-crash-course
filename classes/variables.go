package classes

import "fmt"

func VariablesNotes() {
	// Convensión es CamelCase siendo la primera letra mayuscula para las variables que estan definidas para exportación y con la primera letra minuscula para las variables locales
	var helloVariable string = "Hello world!"
	const constNumber int = 12
	// const ExportableVariable int = 1 Variable exportable

	// El tipado es estático y existen varios tipos por ejemplo de int como int8 int16 int32 etc...

	fmt.Println(helloVariable)
	fmt.Println(constNumber)

	// Casteo de variables
	tempInt := 32
	tempFloat := float64(tempInt)

	println("La temperatura en INT es: ", tempInt)
	println("La temperatura en Float64 es: ", tempFloat)

	fmt.Printf("Tú número %d cebolla %d es: \n", 122, 22)

}
