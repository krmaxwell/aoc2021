package day10

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		got := ValidateChunks("[<(<(<(<{}))><([]([]()")
		assert.Equal(3, got.Score)
		assert.False(got.Valid)
	})

	t.Run("Flag corrupted line with unexpected ']'", func(t *testing.T) {
		got := ValidateChunks("[{[{({}]{}}([{[{{{}}([]")
		assert.Equal(57, got.Score)
		assert.False(got.Valid)
	})

	t.Run("Flag corrupted line with unexpected '}'", func(t *testing.T) {
		got := ValidateChunks("{([(<{}[<>[]}>{[]{[(<()>")
		assert.Equal(1197, got.Score)
		assert.False(got.Valid)
	})

	t.Run("Flag corrupted line with unexpected '>'", func(t *testing.T) {
		got := ValidateChunks("<{([([[(<>()){}]>(<<{{")
		assert.Equal(25137, got.Score)
		assert.False(got.Valid)
	})

	t.Run("Validate complete line", func(t *testing.T) {
		got := ValidateChunks("(<>)")
		assert.True(got.Valid)
	})

	t.Run("Validate incomplete line", func(t *testing.T) {
		got := ValidateChunks(sampleSubsystem[0])
		assert.True(got.Valid)
	})

	t.Run("Sum scores of invalid lines for sample subsystem", func(t *testing.T) {
		sum := 0
		for _, s := range sampleSubsystem {
			got := ValidateChunks(s)
			if !got.Valid {
				sum += got.Score
			}
		}
		assert.Equal(26397, sum)
	})

	t.Run("Sum scores for full subsystem", func(t *testing.T) {
		data := readFullData(t, "day10.dat")
		sum := 0
		for _, s := range data {
			got := ValidateChunks(s)
			if !got.Valid {
				sum += got.Score
			}
		}
		fmt.Println(sum)
	})

	t.Run("Score incomplete line", func(t *testing.T) {
		got := ValidateChunks("[({(<(())[]>[[{[]{<()<>>")
		assert.True(got.Valid)
		assert.Equal(288957, got.Score)
	})

	t.Run("Score another incomplete line", func(t *testing.T) {
		got := ValidateChunks("[(()[<>])]({[<{<<[]>>(")
		assert.True(got.Valid)
		assert.Equal(5566, got.Score)
	})

	t.Run("Find middle score for incomplete sample data", func(t *testing.T) {
		scores := []int{}
		for _, s := range sampleSubsystem {
			got := ValidateChunks(s)
			if got.Valid {
				scores = append(scores, got.Score)
			}
		}
		sort.Ints(scores)
		fmt.Println(scores[len(scores)/2])
	})

	t.Run("Find middle score for incomplete full data", func(t *testing.T) {
		scores := []int{}
		data := readFullData(t, "day10.dat")
		for _, s := range data {
			got := ValidateChunks(s)
			if got.Valid {
				scores = append(scores, got.Score)
			}
		}
		sort.Ints(scores)
		fmt.Println(scores[len(scores)/2])
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
