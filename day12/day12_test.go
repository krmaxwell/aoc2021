package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaveTypeDetection(t *testing.T) {
	assert.True(t, IsSmallCave("b"))
	assert.False(t, IsSmallCave("A"))
}

func TestBuildCaveMap(t *testing.T) {
	input := []string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	}
	caveMap := BuildCaveMap(input)
	assert.Equal(t, 6, len(caveMap))
	assert.Contains(t, caveMap, "start")
	assert.ElementsMatch(t, caveMap["start"], []string{"A", "b"})
	assert.ElementsMatch(t, caveMap["d"], []string{"b"})
	assert.ElementsMatch(t, caveMap["end"], []string{"A", "b"})
}

func Test10PathCave(t *testing.T) {

	// build map of caves
	input := []string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	}
	caveMap := BuildCaveMap(input)
	routes := 0
	FindPaths(caveMap, []string{"start"}, &routes)
	assert.Equal(t, 10, routes)
}

func Test19PathCave(t *testing.T) {
	// build map of caves
	input := []string{
		"dc-end",
		"HN-start",
		"start-kj",
		"dc-start",
		"dc-HN",
		"LN-dc",
		"HN-end",
		"kj-sa",
		"kj-HN",
		"kj-dc",
	}
	caveMap := BuildCaveMap(input)

	routes := 0
	FindPaths(caveMap, []string{"start"}, &routes)
	assert.Equal(t, 19, routes)
}
