package day15

import (
	"fmt"

	"github.com/RyanCarrier/dijkstra"
)

func CreateCavernMap(s []string) [][]int {
	c := make([][]int, len(s))
	for row := range c {
		c[row] = make([]int, len(s[row]))
		for col := range c[row] {
			c[row][col] = int(s[row][col] - '0')

		}
	}
	return c
}

func CreateCavernGraph(cave [][]int) *dijkstra.Graph {
	graph := dijkstra.NewGraph()
	for row := range cave {
		for col := range cave[row] {
			this := fmt.Sprintf("%d,%d", row, col)
			graph.AddMappedVertex(this)
			if row > 0 {
				north := fmt.Sprintf("%d,%d", row-1, col)
				graph.AddMappedArc(this, north, int64(cave[row][col]))
				fmt.Printf("Weight S from %s to %s is %d\n", north, this, cave[row][col])
			}
			if col > 0 {
				west := fmt.Sprintf("%d,%d", row, col-1)
				graph.AddMappedArc(this, west, int64(cave[row][col]))
				fmt.Printf("Weight E from %s to %s is %d\n", west, this, cave[row][col])
			}
		}
	}
	return graph
}

func CavernPath(graph *dijkstra.Graph) int {
	start, err := graph.GetMapping("0,0")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("start is %d\n", start)
	}

	end, err := graph.GetMapping("0,1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("end is %d\n", end)
	}

	path, err := graph.Shortest(start, end)
	if err != nil {
		fmt.Println(err)
	}

	return int(path.Distance)
}
