package main

func Suma(numeros ...int) int {
	resultado := 0
	for _, num := range numeros {
		resultado += num
	}
	return resultado
}

func Resta(a, b int) int {
	return a - b
}
