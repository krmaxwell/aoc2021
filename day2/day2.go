package day2

import (
	"strconv"
	"strings"
)

func PositionTrackNaive(course []string) []int {
	horizontalPos := 0
	depth := 0

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
			depth -= distance
		case "down":
			depth += distance
		}
	}

	return []int{horizontalPos, depth}
}

func PositionTrack(course []string) []int {
	horizontalPos := 0
	depth := 0
	aim := 0

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
			depth += distance * aim
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}

	return []int{horizontalPos, depth}
}
