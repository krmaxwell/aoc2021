package day8

import (
	"strings"
)

func CountSimpleDigits(patterns []string) int {
	var outputValues []string
	var sum int
	for _, s := range patterns {
		nextOutput := strings.Split(s, " | ")[1]
		//fmt.Printf("Found output %q\n", nextOutput)
		outputValues = append(outputValues, nextOutput)
	}

	for _, v := range outputValues {
		for _, x := range strings.Split(v, " ") {
			if len(x) == 2 || len(x) == 4 || len(x) == 3 || len(x) == 7 {
				sum++
			}
		}
	}
	return sum
}
