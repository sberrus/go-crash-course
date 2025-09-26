package main

import (
	"fmt"
)

// FUNCIONES
// Las funciones son bastante standar, la diferencia es que golang no permite la declaración de funciones anidadas. Todas las funciones deben estar a nivel de paquete sino no compila.

// kw <fn name>(    args    ) <return type> {...}
func basicFunc(a int, b int) int {
	return a + b
}

// si pasamos varios argumentos de un mismo tipo, podemos definirlo al final un solo tipo de datos para todos los argumentos.
// En el siguiente ejemplo, todos los argumentos son inferidos como int
func add(a, b, c, d int) int {
	return a + b + c + d
}

// Se pueden hacer grupos también de la siguiente manera.
// los dos primeros argumentos son int, y los otros dos son string.
// Esto se hace con el fin de que sea menos verboso el lenguaje
func addPlusUltra(a, b int, c, d string) int {
	fmt.Println("This is the text given str -> ", c, d)
	return a + b
}

// En go se pueden devolver multiples valores en una función del a siguiente forma
func getFullName(firstName string, secondName string) (string, string) {
	return firstName, secondName
}

func main() {

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

	// FUNCIONES

	// func -> palabra clave para declararla
	// sintax
	// func <func name> (args) <return type> {...}

	fmt.Printf("The function returns: %d\n", basicFunc(1, 2))
	fmt.Printf("The numbers sums: %d\n", add(3, 5, 5, 1))
	fmt.Printf("The numbers sums: %d\n", addPlusUltra(3, 5, "text1", "text2"))

	var firstName, secondName = getFullName("Jhon", "Doe")
	fmt.Println("Hi my name is", firstName, secondName)

	// Una de las cosa que tenemos en go y en otros lenguajes es que si se usa _ como nombre de variable, automaticamente el compilador ignora el valor.
	// Este aproach es realmente útil, porque el compilador fuerza al desarrollador que se usen todas las variables declaradas de forma predeterminada y las que no se usen a la hora de compilar, devuelven error.
	var firstName2, _ = getFullName("Jhonny", "Doe")
	fmt.Println("But they call me", firstName2)

}
