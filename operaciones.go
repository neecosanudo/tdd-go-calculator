package main

// Suma() recibe una cantidad variable de números y devuelve la suma de todos ellos
func Suma(numeros ...int) int {
	// Inicializa el resultado en cero
	resultado := 0

	// Recorre todos los números sumándolos al resultado
	for _, num := range numeros {
		resultado += num
	}

	// Devuelve el resultado
	return resultado
}

// Resta() recibe una cantidad variable de números y devuelve la resta de todos ellos
func Resta(numeros ...int) int {
	// Inicializa la variable resultado con el primer valor de los parámetros
	resultado := numeros[0]

	// Itera sobre cada número en los parámetros
	for _, num := range numeros {
		// Si el número actual no es igual al primer valor, se resta al resultado
		if num != numeros[0] {
			resultado -= num
		}
	}

	// Retorna el resultado final
	return resultado
}
