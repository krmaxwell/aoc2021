package day11

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOctoFlash(t *testing.T) {

	assert := assert.New(t)
	t.Run("flash 1 octopus", func(t *testing.T) {
		octopus := Octopus{Energy: 10, Flash: false}
		octopus.ValidateFlash()
		assert.True(octopus.Flash)
	})

	field := []string{
		"11111",
		"19991",
		"19191",
		"19991",
		"11111",
	}

	t.Run("Construct octopus field", func(t *testing.T) {
		octopi := ConstructOctopi(field)
		assert.Equal(1, octopi[2][2].Energy)
		assert.Equal(9, octopi[1][1].Energy)
	})

	t.Run("Process step 1 for simple field", func(t *testing.T) {
		octopi := ConstructOctopi([]string{
			"000",
			"000",
			"000",
		})
		octopi.ProcessStep()
		want := ConstructOctopi([]string{
			"111",
			"111",
			"111",
		})
		assert.EqualValues(want, octopi)
	})

	t.Run("Process step 1 with 1 flash", func(t *testing.T) {
		octopi := ConstructOctopi([]string{
			"000",
			"090",
			"000",
		})
		flashCount := octopi.ProcessStep()
		want := ConstructOctopi([]string{
			"222",
			"202",
			"222",
		})
		assert.EqualValues(want, octopi)
		assert.Equal(1, flashCount)
	})

	t.Run("Process step 1 with 9 flashes", func(t *testing.T) {
		octopi := ConstructOctopi(field)
		count := octopi.ProcessStep()
		assert.Equal(9, count)
	})

	sampleData := []string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	}

	t.Run("Process sample data 1 step", func(t *testing.T) {
		octopi := ConstructOctopi(sampleData)
		octopi.ProcessStep()
		want := ConstructOctopi([]string{
			"6594254334",
			"3856965822",
			"6375667284",
			"7252447257",
			"7468496589",
			"5278635756",
			"3287952832",
			"7993992245",
			"5957959665",
			"6394862637",
		})
		assert.EqualValues(want, octopi)
	})

	t.Run("Count flashes for 2 steps", func(t *testing.T) {
		octopi := ConstructOctopi(sampleData)
		count := 0
		for i := 0; i < 2; i++ {
			count += octopi.ProcessStep()
		}
		assert.Equal(35, count)
	})

	t.Run("Count flashes for 10 steps", func(t *testing.T) {
		octopi := ConstructOctopi(sampleData)
		count := 0
		for i := 0; i < 10; i++ {
			count += octopi.ProcessStep()
		}
		assert.Equal(204, count)
	})

	t.Run("Count flashes for 100 steps", func(t *testing.T) {
		octopi := ConstructOctopi(sampleData)
		count := 0
		for i := 0; i < 100; i++ {
			count += octopi.ProcessStep()
		}
		assert.Equal(1656, count)
	})

	fullInput := []string{
		"8624818384",
		"3725473343",
		"6618341827",
		"4573826616",
		"8357322142",
		"6846358317",
		"7286886112",
		"8138685117",
		"6161124267",
		"3848415383",
	}
	t.Run("Count flashes in puzzle input", func(t *testing.T) {
		octopi := ConstructOctopi(fullInput)
		count := 0
		for i := 0; i < 100; i++ {
			count += octopi.ProcessStep()
		}
		fmt.Println(count)
	})

	t.Run("Find synchronization in sample data", func(t *testing.T) {
		octopi := ConstructOctopi(sampleData)
		fullFlashCount := len(octopi) * len(octopi[0])
		count := 0
		for i := 0; i < 195; i++ {
			count = octopi.ProcessStep()
		}
		assert.Equal(fullFlashCount, count)
	})

	t.Run("Find synchronization in full data", func(t *testing.T) {
		octopi := ConstructOctopi(fullInput)
		fullFlashCount := len(octopi) * len(octopi[0])
		round := 1
		for {
			if octopi.ProcessStep() == fullFlashCount {
				break
			}
			round++
		}
		fmt.Println(round)
	})
}
