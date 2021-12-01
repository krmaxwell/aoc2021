package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestSonarSweep(t *testing.T) {
	t.Run("naive sweep with test data", func(t *testing.T) {
		reportData := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
		got := SonarSweepNaive(reportData)
		want := 7

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("naive sweep with full data", func(t *testing.T) {
		// this isn't really good TDD practice
		// but that's not the point of this exercise anyway
		reportData := readFullData(t, "day1.dat")
		got := SonarSweepNaive(reportData)
		fmt.Printf("Counted %d decreases\n", got)
	})

	t.Run("sliding sweep with test data", func(t *testing.T) {
		reportData := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
		got := SonarSweepSlidingSum(reportData)
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("sliding sweep with full data", func(t *testing.T) {
		reportData := readFullData(t, "day1.dat")
		got := SonarSweepSlidingSum(reportData)
		fmt.Printf("Counted %d decreases\n", got)
	})
}

func readFullData(t *testing.T, fname string) []int {
	t.Helper()
	f, err := os.Open(fname)
	if err != nil {
		t.Fatalf("Could not open %q: %q", fname, err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	reportData := []int{}
	var datum int

	for scanner.Scan() {
		datum, err = strconv.Atoi(scanner.Text())
		if err != nil {
			t.Fatalf("Could not convert datum: %v", err)
		}
		reportData = append(reportData, datum)
	}

	return reportData
}
