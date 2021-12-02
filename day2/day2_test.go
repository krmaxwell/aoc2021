package day2

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestPositionTrack(t *testing.T) {
	t.Run("naive plotting with example input", func(t *testing.T) {
		course := []string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}
		got := PositionTrackNaive(course)
		want := []int{15, 10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("naive plotting with full input", func(t *testing.T) {
		course := readFullData(t, "day2.dat")
		got := PositionTrackNaive(course)
		fmt.Printf("got %v\n", got)
	})

	t.Run("improved plotting with example input", func(t *testing.T) {
		course := []string{
			"forward 5",
			"down 5",
			"forward 8",
			"up 3",
			"down 8",
			"forward 2",
		}
		got := PositionTrack(course)
		want := []int{15, 60}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("improved plotting with full input", func(t *testing.T) {
		course := readFullData(t, "day2.dat")
		got := PositionTrack(course)
		fmt.Printf("got %v\n", got)
		fmt.Printf("answer %d", got[0]*got[1])
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
