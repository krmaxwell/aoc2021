package day10

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBraces(t *testing.T) {
	sampleSubsystem := []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}
	assert := assert.New(t)

	t.Run("Flag corrupted line with unexpected ')'", func(t *testing.T) {
		assert.Equal(3, TestLine("[<(<(<(<{}))><([]([]()"))
	})

	t.Run("Flag corrupted line with unexpected ']'", func(t *testing.T) {
		assert.Equal(57, TestLine("[{[{({}]{}}([{[{{{}}([]"))
	})

	t.Run("Flag corrupted line with unexpected '}'", func(t *testing.T) {
		assert.Equal(1197, TestLine("{([(<{}[<>[]}>{[]{[(<()>"))
	})

	t.Run("Flag corrupted line with unexpected '>'", func(t *testing.T) {
		assert.Equal(25137, TestLine("<{([([[(<>()){}]>(<<{{"))
	})

	t.Run("Validate complete line", func(t *testing.T) {
		assert.Equal(0, TestLine("(<>)"))
	})

	t.Run("Validate incomplete line", func(t *testing.T) {
		assert.Equal(0, TestLine(sampleSubsystem[0]))
	})

	t.Run("Sum scores for sample subsystem", func(t *testing.T) {
		got := 0
		for _, s := range sampleSubsystem {
			got += TestLine(s)
		}
		assert.Equal(26397, got)
	})

	t.Run("Sum scores for full subsystem", func(t *testing.T) {
		data := readFullData(t, "day10.dat")
		got := 0
		for _, s := range data {
			got += TestLine(s)
		}
		fmt.Println(got)
	})

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
