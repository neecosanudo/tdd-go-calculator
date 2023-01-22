package main

import "testing"

func TestSuma(t *testing.T) {

	t.Run("dos números", func(t *testing.T) {
		got := Suma(2, 5)
		want := 7

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("varios números", func(t *testing.T) {
		got := Suma(4, 5, 6, 10)
		want := 25

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("números negativos", func(t *testing.T) {
		got := Suma(-3, -7, 10)
		want := 0

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
