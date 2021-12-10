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
		assert.Equal(3, ValidateChunks("[<(<(<(<{}))><([]([]()"))
	})

	t.Run("Flag corrupted line with unexpected ']'", func(t *testing.T) {
		assert.Equal(57, ValidateChunks("[{[{({}]{}}([{[{{{}}([]"))
	})

	t.Run("Flag corrupted line with unexpected '}'", func(t *testing.T) {
		assert.Equal(1197, ValidateChunks("{([(<{}[<>[]}>{[]{[(<()>"))
	})

	t.Run("Flag corrupted line with unexpected '>'", func(t *testing.T) {
		assert.Equal(25137, ValidateChunks("<{([([[(<>()){}]>(<<{{"))
	})

	t.Run("Validate complete line", func(t *testing.T) {
		assert.Equal(0, ValidateChunks("(<>)"))
	})

	t.Run("Validate incomplete line", func(t *testing.T) {
		assert.Equal(0, ValidateChunks(sampleSubsystem[0]))
	})

	t.Run("Sum scores for sample subsystem", func(t *testing.T) {
		got := 0
		for _, s := range sampleSubsystem {
			got += ValidateChunks(s)
		}
		assert.Equal(26397, got)
	})

	t.Run("Sum scores for full subsystem", func(t *testing.T) {
		data := readFullData(t, "day10.dat")
		got := 0
		for _, s := range data {
			got += ValidateChunks(s)
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
