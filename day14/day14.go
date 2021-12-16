package day14

import (
	"fmt"
	"math"
	"strings"
)

type Rules map[string][2]string // e.g. {"CH": [2]string{"CB", "HB"}}

type Count map[string]int // e.g. {"NB": 2}

func MakeRules(input []string) Rules {
	rules := Rules{}
	for _, r := range input {
		rs := strings.Split(r, " -> ")
		inFrag := rs[0] // string on LH side of arrow separator
		outFrags := [2]string{
			fmt.Sprintf("%c%s", inFrag[0], rs[1]), // probably a more efficient way exists
			fmt.Sprintf("%s%c", rs[1], inFrag[1]),
		}
		rules[inFrag] = outFrags
	}
	return rules
}

func CountPairs(s string) Count {
	pairs := Count{}
	for i := range s {
		if i == len(s)-1 {
			break
		}
		pairs[s[i:i+2]]++
	}
	return pairs
}

func ProcessPairs(pairs Count, rules Rules, maxSteps int) Count {
	newPairs := Count{}
	for i := 0; i < maxSteps; i++ {
		// each step, go thru each existing pair
		for pair, c := range pairs {
			p0 := rules[pair][0]
			p1 := rules[pair][1]

			newPairs[p0] += c
			newPairs[p1] += c
		}

		pairs = Count{}
		for k, v := range newPairs {
			pairs[k] = v
		}
		newPairs = Count{}
	}
	return pairs
}

func most(pairs Count, final string) int {
	count := Count{}
	//fmt.Println(pairs)
	for p, c := range pairs {
		count[string(p[0])] += c
	}
	count[final]++
	max := 0
	for _, c := range count {
		if c > max {
			max = c
		}
	}
	return max
}

func least(pairs Count, final string) int {
	count := Count{}
	//fmt.Println(pairs)
	for p, c := range pairs {
		count[string(p[0])] += c
	}
	count[final]++
	min := math.MaxInt
	for _, c := range count {
		if c < min {
			min = c
		}
	}
	return min
}
