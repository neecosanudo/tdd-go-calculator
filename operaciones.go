package main

func Suma(numeros ...int) int {
	resultado := 0
	for _, num := range numeros {
		resultado += num
	}
	return resultado
}

func Resta(numeros ...int) int {
	resultado := numeros[0]

	for _, num := range numeros {
		if num != numeros[0] {
			resultado -= num
		}
	}

	return resultado
}
