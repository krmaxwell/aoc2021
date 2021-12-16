package day12

import (
	"strings"
)

// patterned on https://github.com/Rtchaik/AoC-2021/blob/main/Day12/solution.py

type CaveMap map[string][]string

func FindPaths(caveMap map[string][]string, path []string, routes *int) int {
	// implement depth-first search
	// big caves (upper case) can be visited as many times as required on a path
	// small caves (lower case) can only be visited once on a path
	this := path[len(path)-1]
	for _, next := range caveMap[this] {
		if IsLargeCave(next) || !InPath(path, next) {
			if next == "end" {
				//fmt.Println(path)
				*routes++
			} else {
				FindPaths(caveMap, append(path, next), routes)
			}

		}
	}
	return *routes
}

func IsLargeCave(s string) bool {
	return strings.ToUpper(s) == s
}

func InPath(path []string, goal string) bool {
	// O(n) set emulation
	for _, point := range path {
		if point == goal {
			return true
		}
	}
	return false
}

func BuildCaveMap(input []string) CaveMap {
	caves := make(map[string][]string)

	for _, i := range input {
		endpoints := strings.Split(i, "-")
		caves[endpoints[0]] = append(caves[endpoints[0]], endpoints[1])
		caves[endpoints[1]] = append(caves[endpoints[1]], endpoints[0])
	}
	return caves
}
