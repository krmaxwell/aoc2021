package day9

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeightMap(t *testing.T) {

	heightMapStr := []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}
	heights := CreateHeightMap(heightMapStr)

	assert := assert.New(t)

	t.Run("Create height map", func(t *testing.T) {

		hMS := []string{
			"333",
			"333",
			"333",
		}
		h := CreateHeightMap(hMS)
		want := FloorMap{
			{3, 3, 3},
			{3, 3, 3},
			{3, 3, 3},
		}

		assert.Equal(want, h)
	})

	t.Run("Find low points", func(t *testing.T) {
		want := []int{1, 0, 5, 5}
		lowPoints := heights.FindLowPoints()
		lowValues := []int{}
		for i := range lowPoints {
			lowValues = append(lowValues, lowPoints[i].Value)
		}

		assert.EqualValues(want, lowValues)
	})

	t.Run("Check low point center", func(t *testing.T) {
		h := FloorMap{
			{3, 3, 3},
			{3, 0, 3},
			{3, 3, 3},
		}

		assert.True(h.minPoint(1, 1))
	})

	t.Run("Check low point upper row", func(t *testing.T) {
		h := FloorMap{
			{3, 0, 3},
			{3, 3, 3},
			{3, 3, 3},
		}

		assert.True(h.minPoint(0, 1))
	})

	t.Run("Check low point upper left corner", func(t *testing.T) {
		h := FloorMap{
			{0, 3, 3},
			{3, 3, 3},
			{3, 3, 3},
		}

		assert.True(h.minPoint(0, 0))
	})

	t.Run("Check low point upper right corner", func(t *testing.T) {
		h := FloorMap{
			{3, 3, 0},
			{3, 3, 3},
			{3, 3, 3},
		}

		assert.True(h.minPoint(0, 2))
	})

	t.Run("Check not low point", func(t *testing.T) {
		assert.False(heights.minPoint(0, 5))
	})

	t.Run("Find low points in full data", func(t *testing.T) {
		data := readFullData(t, "day9.dat")

		heights := CreateHeightMap(data)
		lowPoints := heights.FindLowPoints()
		sum := 0
		for _, n := range lowPoints {
			sum += n.Value
		}
		sum += len(lowPoints)
		fmt.Println(sum)
	})
}

func TestBasins(t *testing.T) {
	heightMapStr := []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}
	heights := CreateHeightMap(heightMapStr)
	assert := assert.New(t)

	points := heights.FindLowPoints()
	basins := CreateBasinMap(heights, points)
	//basins.Print()

	t.Run("Check point not in basin", func(t *testing.T) {
		assert.Equal(-1, basins[4][0])
	})

	t.Run("Check first low points has basin ID 1", func(t *testing.T) {
		assert.Equal(1, basins[0][0])
	})

	t.Run("Check basin sizes", func(t *testing.T) {
		counts := basins.CountBasins()
		assert.Equal(3, counts[1])
		assert.Equal(9, counts[2])
	})

	t.Run("Create basins from full data", func(t *testing.T) {
		data := readFullData(t, "day9.dat")

		heightMap := CreateHeightMap(data)
		points := heightMap.FindLowPoints()
		basinMap := CreateBasinMap(heightMap, points)
		basinCounts := basinMap.CountBasins()
		basinSizes := make([]int, len(basinCounts)+1)
		for i, b := range basinCounts {
			basinSizes[i] = b
		}
		sort.Ints(basinSizes)
		fmt.Println(basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3])

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
