package day9

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmokeBasin(t *testing.T) {

	heightMapStr := []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}

	assert := assert.New(t)

	t.Run("Create height map", func(t *testing.T) {

		hMS := []string{
			"333",
			"333",
			"333",
		}
		got := CreateMap(hMS)
		want := [][]int{
			{3, 3, 3},
			{3, 3, 3},
			{3, 3, 3},
		}

		assert.Equal(got, want)
	})

	t.Run("Find low points", func(t *testing.T) {
		heightMap := CreateMap(heightMapStr)
		got := FindLowPoints(heightMap)
		want := []int{1, 0, 5, 5}

		assert.EqualValues(got, want)
	})

	t.Run("Check low point center", func(t *testing.T) {
		h := [][]int{
			{3, 3, 3},
			{3, 0, 3},
			{3, 3, 3},
		}
		got := minPoint(h, 1, 1)
		want := true

		assert.Equal(got, want)
	})

	t.Run("Check low point upper row", func(t *testing.T) {
		h := [][]int{
			{3, 0, 3},
			{3, 3, 3},
			{3, 3, 3},
		}
		got := minPoint(h, 0, 1)
		want := true

		assert.Equal(got, want)
	})

	t.Run("Check low point upper left corner", func(t *testing.T) {
		h := [][]int{
			{0, 3, 3},
			{3, 3, 3},
			{3, 3, 3},
		}
		got := minPoint(h, 0, 0)
		want := true

		assert.Equal(got, want)
	})

	t.Run("Check low point upper right corner", func(t *testing.T) {
		h := [][]int{
			{3, 3, 0},
			{3, 3, 3},
			{3, 3, 3},
		}
		got := minPoint(h, 0, 2)
		want := true

		assert.Equal(got, want)
	})

	t.Run("Check not low point", func(t *testing.T) {
		h := CreateMap(heightMapStr)
		got := minPoint(h, 0, 5)
		assert.False(got)
	})

	t.Run("Find low points in full data", func(t *testing.T) {
		data := readFullData(t, "day9.dat")
		h := CreateMap(data)
		lowPoints := FindLowPoints(h)
		sum := 0
		for _, n := range lowPoints {
			sum += n
		}
		sum += len(lowPoints)
		fmt.Println(sum)
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
	data := []string{}

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}
