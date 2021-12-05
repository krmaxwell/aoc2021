package day5

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestMapVents(t *testing.T) {
	exampleInput := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	t.Run("Convert example input to lines", func(t *testing.T) {
		got := ConvertInputToLines(exampleInput)
		want := []Line{
			{x1: 0, y1: 9, x2: 5, y2: 9},
			{x1: 8, y1: 0, x2: 0, y2: 8},
			{x1: 9, y1: 4, x2: 3, y2: 4},
			{x1: 2, y1: 2, x2: 2, y2: 1},
			{x1: 7, y1: 0, x2: 7, y2: 4},
			{x1: 6, y1: 4, x2: 2, y2: 0},
			{x1: 0, y1: 9, x2: 2, y2: 9},
			{x1: 3, y1: 4, x2: 1, y2: 4},
			{x1: 0, y1: 0, x2: 8, y2: 8},
			{x1: 5, y1: 5, x2: 8, y2: 2},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Create diagram from example input", func(t *testing.T) {
		lines := ConvertInputToLines(exampleInput)
		got := CreateDiagramFromLines(lines, 10)
		want := []string{
			".......1..",
			"..1....1..",
			"..1....1..",
			".......1..",
			".112111211",
			"..........",
			"..........",
			"..........",
			"..........",
			"222111....",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ngot  %v \nwant %v", got, want)
		}
	})

	t.Run("Count overlaps from scan", func(t *testing.T) {
		lines := ConvertInputToLines(exampleInput)
		data := ConvertLinesToData(lines, 10)
		got := CountOverlaps(data)
		want := 5
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Count overlaps from full data", func(t *testing.T) {
		input := readFullData(t, "day5.dat")
		lines := ConvertInputToLines(input)
		data := ConvertLinesToData(lines, 1000)
		fmt.Printf("Full data has %d overlaps\n", CountOverlaps(data))
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
	input := []string{}

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}
