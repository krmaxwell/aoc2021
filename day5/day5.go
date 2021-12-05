package day5

import (
	"strconv"
	"strings"
)

type Line struct {
	x1, x2, y1, y2 int64
}

func ConvertInputToLines(input []string) []Line {
	lines := []Line{}
	for _, scanline := range input {
		line := Line{}
		x := strings.Fields(scanline)

		point1 := strings.Split(x[0], ",") // x1,y1
		// note that x[1] is always "->"
		point2 := strings.Split(x[2], ",") // x2,y2

		line.x1, _ = strconv.ParseInt(point1[0], 10, 0)
		line.y1, _ = strconv.ParseInt(point1[1], 10, 0)
		line.x2, _ = strconv.ParseInt(point2[0], 10, 0)
		line.y2, _ = strconv.ParseInt(point2[1], 10, 0)
		lines = append(lines, line)
	}
	return lines
}

func CreateDiagramFromLines(scan []Line, dimension int) []string {
	data := ConvertLinesToData(scan, dimension)
	diagram := make([]string, len(data))
	for row := 0; row < 10; row++ {
		var rowS strings.Builder
		for col := 0; col < 10; col++ {
			if data[col][row] > 0 {
				point := strconv.Itoa(int(data[col][row]))
				rowS.WriteString(point)
			} else {
				rowS.WriteString(".")
			}
		}
		diagram[row] = rowS.String()
	}

	return diagram
}

func ConvertLinesToData(scan []Line, dimension int) [][]int64 {
	var data = make([][]int64, dimension)
	for i := 0; i < dimension; i++ {
		data[i] = make([]int64, dimension)
	}
	for _, line := range scan {
		if line.x1 == line.x2 {
			startY := min(line.y1, line.y2)
			endY := max(line.y1, line.y2)

			for y := startY; y <= endY; y++ {
				data[line.x1][y] += 1
			}

		} else if line.y1 == line.y2 {
			startX := min(line.x1, line.x2)
			endX := max(line.x1, line.x2)

			for x := startX; x <= endX; x++ {
				data[x][line.y1]++
			}
		}
	}
	return data
}

func CountOverlaps(scan [][]int64) int {
	overlaps := 0
	for i := range scan {
		for j := range scan[i] {
			if scan[i][j] > 1 {
				overlaps++
			}
		}
	}
	return overlaps
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
