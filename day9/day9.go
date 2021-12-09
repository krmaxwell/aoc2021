package day9

func CreateMap(strMap []string) [][]int {
	newMap := [][]int{}
	for i := range strMap {
		newLine := []int{}
		for _, c := range strMap[i] {
			newLine = append(newLine, int(c-'0'))
		}
		newMap = append(newMap, newLine)
	}
	return newMap
}

func FindLowPoints(heightMap [][]int) []int {
	lowPoints := []int{}
	for i := range heightMap {
		for j := range heightMap[i] {
			if minPoint(heightMap, i, j) {
				lowPoints = append(lowPoints, heightMap[i][j])
			}
		}
	}
	return lowPoints
}

func minPoint(heightMap [][]int, x, y int) bool {
	point := heightMap[x][y]
	adjacent := []int{}
	min := 9
	if x > 0 {
		adjacent = append(adjacent, heightMap[x-1][y])
	}
	if y > 0 {
		adjacent = append(adjacent, heightMap[x][y-1])
	}
	if x < len(heightMap)-1 {
		adjacent = append(adjacent, heightMap[x+1][y])
	}
	if y < len(heightMap[0])-1 {
		adjacent = append(adjacent, heightMap[x][y+1])
	}
	for _, n := range adjacent {
		if min > n {
			min = n
		}
	}
	return min > point
}
