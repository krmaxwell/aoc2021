package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestSonarSweep(t *testing.T) {
	t.Run("initial test data", func(t *testing.T) {
		reportData := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
		got := SonarSweep(reportData)
		want := 7

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("full data", func(t *testing.T) {
		// this isn't really good TDD practice
		// but that's not the point of this exercise anyway
		f, err := os.Open("day1.dat")
		if err != nil {
			t.Fatalf("Could not open day1.dat: %v", err)
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

		got := SonarSweep(reportData)
		fmt.Printf("Counted %d decreases\n", got)
	})
}
