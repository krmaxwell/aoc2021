package day15

import (
	"testing"

	"github.com/RyanCarrier/dijkstra"
	"github.com/stretchr/testify/assert"
)

func TestSampleCavern(t *testing.T) {
	assert := assert.New(t)
	sampleInput := []string{
		"1163751742",
		"1381373672",
		"2136511328",
		"3694931569",
		"7463417111",
		"1319128137",
		"1359912421",
		"3125421639",
		"1293138521",
		"2311944581",
	}

	t.Run("Create cavern map", func(t *testing.T) {
		cavernMap := CreateCavernMap(sampleInput)
		assert.IsType([][]int{}, cavernMap)
		assert.Len(cavernMap, len(sampleInput))
		assert.Len(cavernMap[0], len(sampleInput[0]))
		assert.Equal(1, cavernMap[0][0])
		assert.Equal(6, cavernMap[0][2])
	})

	t.Run("Create graph from cavern map", func(t *testing.T) {
		cavernMap := CreateCavernMap(sampleInput)
		cavernGraph := CreateCavernGraph(cavernMap)

		assert.IsType(&dijkstra.Graph{}, cavernGraph)
		assert.Len(cavernGraph.Verticies, len(cavernMap)*len(cavernMap[0]))
		assert.Equal(40, CavernPath(cavernGraph))
	})
}
