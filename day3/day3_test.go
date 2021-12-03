package day3

import (
	"bufio"
	"fmt"
	"os"
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

	t.Run("Test most common bit", func(t *testing.T) {
		got := mostCommonBit(3, []string{"10100", "01010", "10000"})
		want := '0'
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Initial test of example data", func(t *testing.T) {
		got := DiagnosticRates(report)
		want := []int{22, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Initial test of full data", func(t *testing.T) {
		report := readFullData(t, "day3.dat")
		got := DiagnosticRates(report)
		fmt.Printf("ℾ %d 𝜀 %d * %d\n", got[0], got[1], got[0]*got[1])
	})
}

func readFullData(t *testing.T, fname string) []string {
	t.Helper()
	f, err := os.Open(fname)
	if err != nil {
		t.Fatalf("Could not open %q: %q", fname, err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	course := []string{}

	for scanner.Scan() {
		course = append(course, scanner.Text())
	}

	return course
}
