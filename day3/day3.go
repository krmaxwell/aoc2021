package day3

import (
	"strconv"
	"strings"
)

func DiagnosticRates(report []string) int {
	var gammaBuilder strings.Builder

	// convertedReport := convertReport(report)
	bitLength := len(report[0])

	for i := 0; i < bitLength; i++ {
		bit := mostCommonBit(i, report)
		gammaBuilder.WriteRune(bit)
	}

	result, _ := strconv.ParseInt(gammaBuilder.String(), 2, 0)
	return int(result)
}

func convertReport(report []string) []int {
	result := []int{}
	for _, number := range report {
		parsed, _ := strconv.ParseInt(number, 2, 0)
		result = append(result, int(parsed))
	}
	return result
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
