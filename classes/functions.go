package classes

import (
	"fmt"
)

// FUNCIONES
// Las funciones son bastante standar, la diferencia es que golang no permite la declaración de funciones anidadas. Todas las funciones deben estar a nivel de paquete sino no compila.

// IMPORTANTE
// Las funciones en go no permiten funciones anidadas, por lo que deben declararse fuera de otras funciones sino el compilador devuelve error.

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

func FunctionsNotes() {
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
