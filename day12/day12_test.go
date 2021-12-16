package day12

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input []string
	want  int
}

func TestCaveTypeDetection(t *testing.T) {
	assert.False(t, IsLargeCave("b"))
	assert.True(t, IsLargeCave("A"))
}

func TestBuildCaveMap(t *testing.T) {
	input := []string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	}
	caveMap := BuildCaveMap(input)
	assert.Equal(t, 6, len(caveMap))
	assert.Contains(t, caveMap, "start")
	assert.ElementsMatch(t, caveMap["start"], []string{"A", "b"})
	assert.ElementsMatch(t, caveMap["d"], []string{"b"})
	assert.ElementsMatch(t, caveMap["end"], []string{"A", "b"})
}

func TestPart1(t *testing.T) {
	testCases := []testCase{
		{input: []string{
			"start-A",
			"start-b",
			"A-c",
			"A-b",
			"b-d",
			"A-end",
			"b-end",
		},
			want: 10,
		},
		{
			input: []string{
				"dc-end",
				"HN-start",
				"start-kj",
				"dc-start",
				"dc-HN",
				"LN-dc",
				"HN-end",
				"kj-sa",
				"kj-HN",
				"kj-dc",
			},
			want: 19,
		},
		{
			input: []string{
				"fs-end",
				"he-DX",
				"fs-he",
				"start-DX",
				"pj-DX",
				"end-zg",
				"zg-sl",
				"zg-pj",
				"pj-he",
				"RW-he",
				"fs-DX",
				"pj-RW",
				"zg-RW",
				"start-pj",
				"he-WI",
				"zg-he",
				"pj-fs",
				"start-RW",
			},
			want: 226,
		},
	}
	for _, c := range testCases {
		caveMap := BuildCaveMap(c.input)
		routes := 0
		FindPaths(caveMap, []string{"start"}, &routes)
		assert.Equal(t, c.want, routes)
	}
}

func TestFullInputPart1(t *testing.T) {
	input := []string{
		"lg-GW",
		"pt-start",
		"pt-uq",
		"nx-lg",
		"ve-GW",
		"start-nx",
		"GW-start",
		"GW-nx",
		"pt-SM",
		"sx-GW",
		"lg-end",
		"nx-SM",
		"lg-SM",
		"pt-nx",
		"end-ve",
		"ve-SM",
		"TG-uq",
		"end-SM",
		"SM-uq",
	}

	caveMap := BuildCaveMap(input)

	routes := 0
	FindPaths(caveMap, []string{"start"}, &routes)
	fmt.Println(routes)
}
