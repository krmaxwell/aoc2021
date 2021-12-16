package day14

import (
	"fmt"
	"strings"
)

func MakeRules(input []string) map[string]string {
	rules := make(map[string]string, len(input))
	for _, r := range input {
		rs := strings.Split(r, " -> ")
		rules[rs[0]] = rs[1]
	}
	return rules
}

func MakePairs(s string) map[string]int {
	pairs := make(map[string]int)
	for i := range s {
		if i == len(s)-1 {
			break
		}
		pairs[s[i:i+2]]++
	}
	return pairs
}

func ProcessPairs(pairs map[string]int, rules map[string]string) map[string]int {
	newPairs := make(map[string]int)
	for p := range pairs {
		p0 := fmt.Sprintf("%c%s", p[0], rules[p])
		p1 := fmt.Sprintf("%s%c", rules[p], p[1])
		newPairs[p0]++
		newPairs[p1]++
	}
	return newPairs
}

func most(pairs map[string]int, final string) int {
	count := make(map[byte]int)
	//fmt.Println(pairs)
	for p := range pairs {
		count[p[0]] += pairs[p]
	}
	count[final[0]]++
	max := 0
	for p := range count {
		//fmt.Printf("%c: %d\n", p, c[p])
		if count[p] > max {
			max = count[p]
		}
	}
	return max
}
