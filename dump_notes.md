# Go AFAP Crash Course

- Instalar binario y configurarlo https://go.dev/doc/install

## Tabla de contenidos
- [Go AFAP Crash Course](#go-afap-crash-course)
  - [Tabla de contenidos](#tabla-de-contenidos)
  - [Modulos](#modulos)
  - [Packages](#packages)
  - [Ejecución de código](#ejecución-de-código)
  - [Descarga de dependencias.](#descarga-de-dependencias)
  - [Tiempo de ejecución](#tiempo-de-ejecución)


## Modulos
Una de las cosas que hay que hacer en principio es la de incializar el proyecto, esto lo hacemos con el comando `go mod init <nombre de repo>`. Como estandar, se recomienda poner como nombre del package el repositorio donde se va a almacenar el codigo fuente y luego el nombre del paquete ej: github.com/<nombre paquete>. 

Esto se hace con fines de poder cumplir con los requisitos de el manejo de dependencias del lenguaje.

Con esto, se crea el fichero go.mod y a partir de este se controlan las dependencias que vayamos utilizando en nuestro proyecto.

## Packages
Los packages son una forma lógica de organizar las funciones que pertenecen a un módulo. 

Automaticamente, go lo que hace es que a partir de un directorio, busca los diferentes paquetes y los corelaciona para poder realizar los imports de forma ordenada. 

Se puede evidenciar esto cuando tienes dos funciones main en el mismo directorio. 

Esto hace que el compilador se queje porque esta detectando múltiples funciones main qeu son las que usa go para ejecutar de primero cuando estas en un codigo fuente.

## Ejecución de código
`go run <ruta paquete>`

## Descarga de dependencias.
En go, las dependencias y repositorios de la comunidad los puedes buscar en https://pkg.go.dev/ siendo este le buscador de dependencias de la comunidad.

A diferencia de otros lenguajes, en go, hay un comando `go mod tidy` que lo que hace es encontrar en todo el código los `imports` y en el caso de que encuentre uno que no este instalado como dependencia, go lo instala.

Puede pasar que pasado un tiempo, ya no uses una, al ejecutar el comando `go mod tidy` go automaticamente se encarga de eliminar los binarios de la dependencia que no este siendo utilizada.

## Tiempo de ejecución
A diferencia de sus competidores, Go entra en el grupo de lenguajes compilados.

Siendo éste uno de los grupos más rápidos en tiempo de ejecución debido a ser compilado. En el siguiente enlace [stackoverflow (Interpreters vs Compilers vs Virtual Machines)](https://stackoverflow.com/questions/14678060/interpreters-vs-compilers-vs-virtual-machines), hay una discusión que profundiza más en este asunto.

Dicho lo anterior, golang es rápido.

Una de las formas que tiene golang de ir rápido a la vez de ser seguro, es la posibilidad de tener máquinas virtuales, pero de una forma más eficiente y rápida. Sería como una miniaturización de como trabajan las VM de Java o C#.