package day2

import (
	"strconv"
	"strings"
)

func PositionTrack(course []string) []int {
	horizontalPos := 0
	verticalPos := 0
	var currentInput []string
	var direction string
	var distance int
	for _, next := range course {
		currentInput = strings.Split(next, " ")
		direction = currentInput[0]
		distance, _ = strconv.Atoi(currentInput[1])
		switch direction {
		case "forward":
			horizontalPos += distance
		case "up":
			verticalPos -= distance
		case "down":
			verticalPos += distance
		}
	}
	return []int{horizontalPos, verticalPos}
}
