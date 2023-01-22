package main

import "testing"

func TestSuma(t *testing.T) {

	t.Run("dos números", func(t *testing.T) {
		got := Suma(2, 5)
		want := 7

		assertMessageForSumaTest(t, got, want)
	})

	t.Run("varios números", func(t *testing.T) {
		got := Suma(4, 5, 6, 10)
		want := 25

		assertMessageForSumaTest(t, got, want)
	})

	t.Run("números negativos", func(t *testing.T) {
		got := Suma(-3, -7, 10)
		want := 0

		assertMessageForSumaTest(t, got, want)
	})
}

func TestResta(t *testing.T) {
	t.Run("dos números", func(t *testing.T) {
		got := Resta(15, 4)
		want := 11

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func assertMessageForSumaTest(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
