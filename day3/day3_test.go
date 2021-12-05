package day3

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestPowerConsumptionRates(t *testing.T) {
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
		got := commonBit(3, []string{"10100", "01010", "10000"}, 1)
		want := '0'
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Power consumption test of example data", func(t *testing.T) {
		got := PowerConsumptionRates(report)
		want := Rates{Gamma: 22, Epsilon: 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Power consumption test of full data", func(t *testing.T) {
		report := readFullData(t, "day3.dat")
		got := PowerConsumptionRates(report)
		fmt.Printf("𝝲 %d 𝞊 %d * %d\n", got.Gamma, got.Epsilon, got.Gamma*got.Epsilon)
	})
}

func TestLifeSupportRating(t *testing.T) {
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

	t.Run("Oxygen rating in sample data", func(t *testing.T) {
		got := OxygenRating(report)
		want := int64(23)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("CO2 rating in sample data", func(t *testing.T) {
		got := CO2Rating(report)
		want := int64(10)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Life support rating for full data", func(t *testing.T) {
		report := readFullData(t, "day3.dat")
		oxy := OxygenRating(report)
		co2 := CO2Rating(report)
		fmt.Printf("O2 %d CO2 %d * %d\n", oxy, co2, oxy*co2)
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
