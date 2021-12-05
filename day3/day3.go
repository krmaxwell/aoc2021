package day3

import (
	"fmt"
	"strconv"
	"strings"
)

type Rates struct {
	Gamma   int64
	Epsilon int64
}

func PowerConsumptionRates(report []string) Rates {
	var gammaBuilder strings.Builder
	var epsilonBuilder strings.Builder

	bitLength := len(report[0])

	for i := 0; i < bitLength; i++ {
		mostCommonBit := commonBit(i, report, 1)
		leastCommonBit := commonBit(i, report, 0)
		gammaBuilder.WriteRune(mostCommonBit)
		epsilonBuilder.WriteRune(leastCommonBit)
	}

	// note that we only return '0' or '1' so ParseInt() is guaranteed to succeed
	gamma, _ := strconv.ParseInt(gammaBuilder.String(), 2, 0)
	epsilon, _ := strconv.ParseInt(epsilonBuilder.String(), 2, 0)

	return Rates{Gamma: gamma, Epsilon: epsilon}
}

func OxygenRating(report []string) int64 {
	bitLength := len(report[0])

	for i := 0; i < bitLength; i++ {
		mostCommonBit := commonBit(i, report, 1)
		newReport := []string{}
		for _, reading := range report {
			if reading[i] == byte(mostCommonBit) {
				newReport = append(newReport, reading)
			}
		}
		report = newReport
	}
	if len(report) > 1 {
		panic(fmt.Sprintf("Didn't filter to 1 result (%d)", len(report)))
	}
	result, _ := strconv.ParseInt(report[0], 2, 0)
	return result
}

func CO2Rating(report []string) int64 {
	bitLength := len(report[0])

	for i := 0; i < bitLength; i++ {
		leastCommonBit := commonBit(i, report, 0)
		fmt.Printf("Least common bit for pos %d is %q\n", i, leastCommonBit)
		newReport := []string{}
		for _, reading := range report {
			if reading[i] == byte(leastCommonBit) {
				newReport = append(newReport, reading)
			}
		}
		report = newReport
		if len(report) == 1 {
			break
		}
	}
	if len(report) > 1 {
		panic(fmt.Sprintf("Didn't filter to 1 result (%d)", len(report)))
	}
	result, _ := strconv.ParseInt(report[0], 2, 0)
	return result
}

func commonBit(pos int, report []string, flag byte) rune {
	zeroes := 0
	ones := 0

	for _, binary := range report {
		if binary[pos] == '0' {
			zeroes++
		} else {
			ones++
		}
	}

	if flag == 1 && ones >= zeroes {
		return '1'
	} else if flag == 0 && zeroes > ones {
		return '1'
	}
	return '0'

}
