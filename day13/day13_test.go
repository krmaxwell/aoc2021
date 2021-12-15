package day13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFold(t *testing.T) {
	assert := assert.New(t)

	t.Run("Test paper creation", func(t *testing.T) {
		paper := MakePaper(10, 14)
		assert.Equal(10, len(paper[0]))
		assert.Equal(14, len(paper))
	})

	sampleData := []Point{
		{6, 10},
		{0, 14},
		{9, 10},
		{0, 3},
		{10, 4},
		{4, 11},
		{6, 0},
		{6, 12},
		{4, 1},
		{0, 13},
		{10, 12},
		{3, 4},
		{3, 0},
		{8, 4},
		{1, 10},
		{2, 14},
		{8, 10},
		{9, 0},
	}

	t.Run("Set dots from sample data", func(t *testing.T) {
		paper := MakePaper(10+1, 14+1) // TODO: auto-find max X and Y
		paper.SetDots(sampleData)

		assert.Equal(1, paper.CheckDot(6, 10))
		assert.Equal(1, paper.CheckDot(10, 12))
		assert.Equal(0, paper.CheckDot(1, 4))
	})

	t.Run("Test 1 fold on sample data", func(t *testing.T) {
		paper := MakePaper(10+1, 14+1) // TODO: auto-find max X and Y
		paper.SetDots(sampleData)
		paper = paper.Fold(0, 7)
		assert.Equal(17, paper.CountDots())
		paper = paper.Fold(5, 0)
		assert.Equal(16, paper.CountDots())
		paper.Print()
	})

	t.Run("Fold full data", func(t *testing.T) {
		points := readFullData(t, "day13.dat")
		mX := 0
		mY := 0
		for _, p := range points {
			if p.X > mX {
				mX = p.X
			}
			if p.Y > mY {
				mY = p.Y
			}
		}
		paper := MakePaper(mX+1, mY+1)
		paper.SetDots(points)
		paper = paper.Fold(655, 0)
		fmt.Println(paper.CountDots())
		paper = paper.Fold(0, 447)
		paper = paper.Fold(327, 0)
		paper = paper.Fold(0, 223)
		paper = paper.Fold(163, 0)
		paper = paper.Fold(0, 111)
		paper = paper.Fold(81, 0)
		paper = paper.Fold(0, 55)
		paper = paper.Fold(40, 0)
		paper = paper.Fold(0, 27)
		paper = paper.Fold(0, 13)
		paper = paper.Fold(0, 6)
		paper.Print()
	})

}

func readFullData(t *testing.T, fname string) []Point {
	t.Helper()
	f, err := os.Open(fname)
	if err != nil {
		t.Fatalf("Could not open %q: %q", fname, err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	data := []Point{}

	for scanner.Scan() {
		s := scanner.Text()
		pStr := strings.Split(s, ",")
		pX, _ := strconv.Atoi(pStr[0])
		pY, _ := strconv.Atoi(pStr[1])
		data = append(data, Point{pX, pY})
	}

	return data
}
