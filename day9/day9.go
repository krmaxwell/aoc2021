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

func (f FloorMap) InBasin(x, y int) bool {
	return f[x][y] < 9
}
