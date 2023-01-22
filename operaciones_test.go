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
}
