package day6

import (
	"fmt"
	"testing"
)

func TestLanternfish(t *testing.T) {
	t.Run("simulate sample data short run", func(t *testing.T) {
		fish := []int{3, 4, 3, 1, 2}
		got := simulateDays(fish, 18)
		want := 26

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("simulate sample data long run", func(t *testing.T) {
		fish := []int{3, 4, 3, 1, 2}
		got := simulateDays(fish, 80)
		want := 5934

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("simulate full data long run", func(t *testing.T) {
		fish := []int{2, 5, 3, 4, 4, 5, 3, 2, 3, 3, 2, 2, 4, 2, 5, 4, 1, 1, 4, 4, 5, 1, 2, 1, 5, 2, 1, 5, 1, 1, 1, 2, 4, 3, 3, 1, 4, 2, 3, 4, 5, 1, 2, 5, 1, 2, 2, 5, 2, 4, 4, 1, 4, 5, 4, 2, 1, 5, 5, 3, 2, 1, 3, 2, 1, 4, 2, 5, 5, 5, 2, 3, 3, 5, 1, 1, 5, 3, 4, 2, 1, 4, 4, 5, 4, 5, 3, 1, 4, 5, 1, 5, 3, 5, 4, 4, 4, 1, 4, 2, 2, 2, 5, 4, 3, 1, 4, 4, 3, 4, 2, 1, 1, 5, 3, 3, 2, 5, 3, 1, 2, 2, 4, 1, 4, 1, 5, 1, 1, 2, 5, 2, 2, 5, 2, 4, 4, 3, 4, 1, 3, 3, 5, 4, 5, 4, 5, 5, 5, 5, 5, 4, 4, 5, 3, 4, 3, 3, 1, 1, 5, 2, 4, 5, 5, 1, 5, 2, 4, 5, 4, 2, 4, 4, 4, 2, 2, 2, 2, 2, 3, 5, 3, 1, 1, 2, 1, 1, 5, 1, 4, 3, 4, 2, 5, 3, 4, 4, 3, 5, 5, 5, 4, 1, 3, 4, 4, 2, 2, 1, 4, 1, 2, 1, 2, 1, 5, 5, 3, 4, 1, 3, 2, 1, 4, 5, 1, 5, 5, 1, 2, 3, 4, 2, 1, 4, 1, 4, 2, 3, 3, 2, 4, 1, 4, 1, 4, 4, 1, 5, 3, 1, 5, 2, 1, 1, 2, 3, 3, 2, 4, 1, 2, 1, 5, 1, 1, 2, 1, 2, 1, 2, 4, 5, 3, 5, 5, 1, 3, 4, 1, 1, 3, 3, 2, 2, 4, 3, 1, 1, 2, 4, 1, 1, 1, 5, 4, 2, 4, 3}
		fmt.Println(simulateDays(fish, 80))
	})
}
