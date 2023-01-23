package main

import "fmt"

func main() {

	// Declara variables con valores constantes
	const (
		a = 20
		b = -7
		c = 15
		d = 89
	)

	/* Asigna a la variable resultadoSuma el resultado de la función Suma() y
	a la variable resultadoResta el resultado de la función Resta(),
	utilizando las variables a, b, c y d como parámetros
	*/
	resultadoSuma := Suma(a, b, c, d)
	resultadoResta := Resta(a, b, c, d)

	/* Imprime en consola el resultado de las operaciones de suma y resta
	utilizando las variables a, b, c y d como parámetros
	*/
	fmt.Printf("Dado los números %d, %d, %d y %d:\nLa suma de todos es: %d.\nLa resta de todos es: %d.\n", a, b, c, d, resultadoSuma, resultadoResta)
}
