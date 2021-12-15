package day9

import "fmt"

type FloorMap [][]int

type Point struct {
	Value, X, Y int
}

func CreateHeightMap(strMap []string) FloorMap {
	var newMap FloorMap
	for i := range strMap {
		newLine := []int{}
		for _, c := range strMap[i] {
			newLine = append(newLine, int(c-'0'))
		}
		newMap = append(newMap, newLine)
	}
	return newMap
}

func CreateBasinMap(heights FloorMap, lowPoints []Point) FloorMap {

	// initialization
	basins := make([][]int, len(heights))
	for i := range heights {
		basins[i] = make([]int, len(heights[0]))
		for j := range basins[i] {
			if heights[i][j] == 9 { // ridges are not in any basin
				basins[i][j] = -1
			} else {
				basins[i][j] = 0 // we'll figure out which basin later
			}
		}
	}

	// flood fill
	basinID := 1
	for _, p := range lowPoints {
		//fmt.Printf("Flood fill basin %d starting at x %d y %d\n", basinID, p.X, p.Y)
		basins = floodFill(basins, p.X, p.Y, basinID)
		basinID++
	}
	return basins
}

func floodFill(basinMap [][]int, x, y, basinID int) [][]int {
	if x < 0 || y < 0 || x >= len(basinMap) || y >= len(basinMap[0]) { // point outside the map
		return basinMap
	}
	if basinMap[x][y] == -1 { // ridge
		return basinMap
	}
	if basinMap[x][y] == 0 {
		//fmt.Printf("Set x %d y %d to %d\n", x, y, basinID)
		basinMap[x][y] = basinID
	} else {
		//fmt.Printf("Found x %d y %d is already %d\n", x, y, basinID)
		return basinMap
	}
	basinMap = floodFill(basinMap, x+1, y, basinID) // south
	basinMap = floodFill(basinMap, x-1, y, basinID) // north
	basinMap = floodFill(basinMap, x, y+1, basinID) // east
	basinMap = floodFill(basinMap, x, y-1, basinID) // west
	return basinMap
}

func (f FloorMap) FindLowPoints() []Point {
	lowPoints := []Point{}
	for i := range f {
		for j := range f[i] {
			if f.minPoint(i, j) {
				point := Point{Value: f[i][j], X: i, Y: j}
				lowPoints = append(lowPoints, point)
			}
		}
	}
	return lowPoints
}

func (f FloorMap) minPoint(x, y int) bool {
	point := f[x][y]
	adjacent := []int{}
	min := 9
	if x > 0 {
		adjacent = append(adjacent, f[x-1][y])
	}
	if y > 0 {
		adjacent = append(adjacent, f[x][y-1])
	}
	if x < len(f)-1 {
		adjacent = append(adjacent, f[x+1][y])
	}
	if y < len(f[0])-1 {
		adjacent = append(adjacent, f[x][y+1])
	}
	for _, n := range adjacent {
		if min > n {
			min = n
		}
	}
	return min > point
}

func (f FloorMap) Print() {
	for i := range f {
		for j := range f[0] {
			if f[i][j] >= 0 {
				fmt.Print(f[i][j])
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func (f FloorMap) CountBasins() map[int]int {
	var basins = make(map[int]int)
	for i := range f {
		for j := range f[0] {
			if f[i][j] < 0 { // don't count ridges
				continue
			}
			basins[f[i][j]]++
		}
	}
	return basins
}
