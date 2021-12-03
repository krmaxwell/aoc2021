package day3

import (
	"reflect"
	"testing"
)

func TestBinaryDiagnostic(t *testing.T) {
	report := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	t.Run("Test binary conversion", func(t *testing.T) {
		got := convertReport([]string{"00100", "01010"})
		want := []int{4, 10}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Test most common bit", func(t *testing.T) {
		got := mostCommonBit(3, []string{"10100", "01010", "10000"})
		want := '0'
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Initial test", func(t *testing.T) {
		got := DiagnosticRates(report)
		want := 22

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
