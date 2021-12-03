package day3

import (
	"strconv"
	"strings"
)

func DiagnosticRates(report []string) []int {
	var gammaBuilder strings.Builder
	var epsilonBuilder strings.Builder

	bitLength := len(report[0])

	for i := 0; i < bitLength; i++ {
		commonBit := mostCommonBit(i, report)
		leastBit := leastCommonBit(i, report)
		gammaBuilder.WriteRune(commonBit)
		epsilonBuilder.WriteRune(leastBit)
	}

	gamma, _ := strconv.ParseInt(gammaBuilder.String(), 2, 0)
	epsilon, _ := strconv.ParseInt(epsilonBuilder.String(), 2, 0)
	return []int{int(gamma), int(epsilon)}
}

func mostCommonBit(pos int, report []string) rune {
	zeroes := 0
	ones := 0

	for _, binary := range report {
		if binary[pos] == '0' {
			zeroes++
		} else {
			ones++
		}
	}

	if zeroes > ones {
		return '0'
	}
	return '1'
}

func leastCommonBit(pos int, report []string) rune {
	zeroes := 0
	ones := 0

	for _, binary := range report {
		if binary[pos] == '0' {
			zeroes++
		} else {
			ones++
		}
	}

	if zeroes < ones {
		return '0'
	}
	return '1'
}
