package day8

import (
	"sort"
	"strconv"
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

func FindNumberZero(words []string, fourDiff string) string {
	for _, w := range words {
		if len(w) == 6 && !contains(w, fourDiff) {
			return SortString(w)
		}
	}
	return ""
}

func FindNumberOne(words []string) string {
	for _, w := range words {
		if len(w) == 2 {
			return SortString(w)
		}
	}
	return ""
}

func FindNumberTwo(words []string, three, five string) string {
	for _, w := range words {
		if len(w) == 5 && !contains(w, three) && !contains(w, five) {
			return SortString(w)
		}
	}
	return ""
}

// five letters, includes all letters from one
func FindNumberThree(words []string, one string) string {
	for _, w := range words {
		if len(w) == 5 && contains(w, one) {
			return SortString(w)
		}
	}
	return ""
}

func FindNumberFour(words []string) string {
	for _, w := range words {
		if len(w) == 4 {
			return SortString(w)
		}
	}
	return ""
}

func FindNumberFive(words []string, fourDiff string) string {
	for _, w := range words {
		if len(w) == 5 && contains(w, fourDiff) {
			return SortString(w)
		}
	}
	return ""
}

func FindNumberSix(words []string, four, fourDiff string) string {
	for _, w := range words {
		if len(w) == 6 && contains(w, fourDiff) && !contains(w, four) {
			return SortString(w)
		}
	}
	return ""
}

func FindNumberSeven(words []string) string {
	for _, w := range words {
		if len(w) == 3 {
			return SortString(w)
		}
	}
	return ""
}

func FindNumberNine(words []string, four string) string {
	for _, w := range words {
		if len(w) == 6 && contains(w, four) {
			return SortString(w)
		}
	}
	return ""
}

// https://stackoverflow.com/a/22689818/1569808
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func DecodeSignal(signal string) [10]string {
	words := strings.Split(signal, " ")

	// basic logic from https://www.reddit.com/r/adventofcode/comments/rbvpui/2021_day_8_part_2_my_logic_on_paper_i_used_python/
	key := [10]string{}
	key[8] = "abcdefg" // always will have all 7 segments
	key[1] = FindNumberOne(words)
	key[4] = FindNumberFour(words)
	fourDiff := wordsDiff(key[4], key[1])
	key[7] = FindNumberSeven(words)
	key[3] = FindNumberThree(words, key[1])
	key[5] = FindNumberFive(words, fourDiff)
	key[2] = FindNumberTwo(words, key[3], key[5])
	key[9] = FindNumberNine(words, key[4])
	key[6] = FindNumberSix(words, key[4], fourDiff)
	key[0] = FindNumberZero(words, fourDiff)
	return key
}

func OutputValue(signal string, key [10]string) int {
	outputS := strings.Split(signal, "|")[1]
	outputWords := strings.Split(outputS, " ")
	var valueS strings.Builder
	for _, w := range outputWords {
		for i, v := range key {
			if SortString(w) == v {
				//fmt.Printf("%q %q %d\n", w, v, i)
				valueS.WriteString(strconv.Itoa(i))
			}
		}
	}

	value, _ := strconv.Atoi(valueS.String())
	return value
}

// find letters in s1 not in s2
func wordsDiff(s1, s2 string) string {
	var diff strings.Builder

	for _, r := range s1 {
		if !strings.ContainsRune(s2, r) {
			diff.WriteRune(r)
		}
	}
	return SortString(diff.String())
}

func contains(s1, s2 string) bool {
	for _, r := range s2 {
		if !strings.ContainsRune(s1, r) {
			return false
		}
	}
	return true
}
