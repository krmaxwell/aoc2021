package day14

import (
	"bufio"
	"fmt"
	"os"
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

	counts := CountPairs(template)
	assert.Contains(counts, "NN")
	assert.Contains(counts, "NC")
	assert.Contains(counts, "CB")
	fmt.Println(counts)

	counts = ProcessPairs(counts, rules, 10)

	mostC := most(counts, "B")
	assert.Equal(1749, mostC)

	leastC := least(counts, "B")
	assert.Equal(161, leastC)
	fmt.Println(mostC - leastC)
}

func TestFullData(t *testing.T) {
	template := "SFBBNKKOHHHPFOFFSPFV"

	input := readFullData(t, "day14.dat")
	rules := MakeRules(input)
	counts := CountPairs(template)
	counts = ProcessPairs(counts, rules, 10)

	mostC := most(counts, "V")
	leastC := least(counts, "V")
	fmt.Println(mostC - leastC)

}

func readFullData(t *testing.T, fname string) []string {
	t.Helper()
	f, err := os.Open(fname)
	if err != nil {
		t.Fatalf("Could not open %q: %q", fname, err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	data := []string{}

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}
