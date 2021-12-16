package day14

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleData(t *testing.T) {
	assert := assert.New(t)
	template := "NNCB"

	ruleInput := []string{
		"CH -> B",
		"HH -> N",
		"CB -> H",
		"NH -> C",
		"HB -> C",
		"HC -> B",
		"HN -> C",
		"NN -> C",
		"BH -> H",
		"NC -> B",
		"NB -> B",
		"BN -> B",
		"BB -> N",
		"BC -> B",
		"CC -> N",
		"CN -> C",
	}

	rules := MakeRules(ruleInput)
	assert.Len(rules, len(ruleInput))
	assert.Contains(rules, "CH")

	pairs := MakePairs(template)
	assert.Len(pairs, 3)
	assert.Contains(pairs, "NN")
	assert.Contains(pairs, "NC")
	assert.Contains(pairs, "CB")

	pairs = ProcessPairs(pairs, rules)

	// after 1 step
	assert.Equal(2, most(pairs, "B"))

	pairs = ProcessPairs(pairs, rules)

	// 39 more steps
	for i := 0; i < 39; i++ {
		pairs = ProcessPairs(pairs, rules)
	}
	fmt.Println(pairs)
	assert.Equal(1749, most(pairs, "B"))
}
