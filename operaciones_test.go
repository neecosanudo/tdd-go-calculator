package main

import "testing"

func TestSuma(t *testing.T) {

	t.Run("Suma de dos números", func(t *testing.T) {
		// Se realiza la suma de los números 2 y 5
		got := Suma(2, 5)
		// Se espera que el resultado sea 7
		want := 7

		// Se utiliza la función assertMessage para comparar el resultado obtenido y el esperado
		assertMessage(t, got, want)
	})

	t.Run("Suma de varios números", func(t *testing.T) {
		// Se realiza la suma de los números 4, 5, 6 y 10
		got := Suma(4, 5, 6, 10)
		// Se espera que el resultado sea 25
		want := 25

		// Se utiliza la función assertMessage para comparar el resultado obtenido y el esperado
		assertMessage(t, got, want)
	})

	t.Run("Suma de números negativos", func(t *testing.T) {
		// Se realiza la suma de los números -3, -7 y 10
		got := Suma(-3, -7, 10)
		// Se espera que el resultado sea 0
		want := 0

		// Se utiliza la función assertMessage para comparar el resultado obtenido y el esperado
		assertMessage(t, got, want)
	})
}

func TestResta(t *testing.T) {
	// Prueba la función Resta con dos números
	t.Run("Resta de dos números", func(t *testing.T) {
		// Asigna el resultado de la función Resta a la variable "got"
		got := Resta(15, 4)
		// Se espera que el resultado sea 11
		want := 11

		// Compara el resultado esperado con el resultado obtenido
		assertMessage(t, got, want)
	})

	// Prueba la función Resta con varios números
	t.Run("Resta de varios números", func(t *testing.T) {
		// Asigna el resultado de la función Resta a la variable "got"
		got := Resta(100, 20, 50, 2, 2, 1)
		// Se espera que el resultado sea 25
		want := 25

		// Compara el resultado esperado con el resultado obtenido
		assertMessage(t, got, want)
	})
}

func assertMessage(t *testing.T, got, want int) {
	// La función t.Helper() indica que esta función es una función auxiliar y no debe ser considerada al calcular el número de fallos
	t.Helper()

	// Compara el valor obtenido (got) con el valor esperado (want)
	if got != want {
		// Si son diferentes, imprime un mensaje de error con el valor obtenido y el esperado
		t.Errorf("got %d, want %d", got, want)
	}
}
