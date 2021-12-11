package day9

type FloorMap [][]int

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

func CreateBasinMap(heights FloorMap) FloorMap {
	basins := make([][]int, len(heights))
	for i := range heights {
		basins[i] = make([]int, len(heights[i]))
	}
	currentBasin := 0

	for i := range basins {
		for j := range basins[i] {
			if heights[i][j] == 9 {
				basins[i][j] = -1 // top of a ridge is not in a basin
			} else {
				if i > 0 && basins[i-1][j] > -1 {
					// match the basin ID above us
					basins[i][j] = basins[i-1][j]
				} else if j > 0 && basins[i][j-1] > -1 {
					// match the basin ID to our left
					basins[i][j] = basins[i][j-1]
				} else if i < len(basins)-1 && basins[i+1][j] > -1 {
					// match the basin ID below us
					basins[i][j] = basins[i+1][j]
				} else if j < len(basins[i])-1 && basins[i][j+1] > -1 {
					// match the basin ID to our right
					basins[i][j] = basins[i][j+1]
				} else {
					// new basin
					basins[i][j] = currentBasin + 1
					currentBasin++
				}

			}
		}
	}
	return basins
}

func (f FloorMap) FindLowPoints() []int {
	lowPoints := []int{}
	for i := range f {
		for j := range f[i] {
			if f.minPoint(i, j) {
				lowPoints = append(lowPoints, f[i][j])
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
