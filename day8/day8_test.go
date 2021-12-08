package day8

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestSignalPatterns(t *testing.T) {

	sampleSignalPattern := []string{
		"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
		"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
		"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
		"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
		"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
		"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
		"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
		"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
		"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
		"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
	}

	t.Run("Count output 1", func(t *testing.T) {
		signalPattern1 := []string{"xxx | fcgedb cgadb ebacf gc"}
		got := CountSimpleDigits(signalPattern1)
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Count output 4", func(t *testing.T) {
		signalPattern4 := []string{
			"xxx | fdgac cefdb cefbgd gcbe"}
		got := CountSimpleDigits(signalPattern4)
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Count output 7", func(t *testing.T) {
		signalPattern7 := []string{
			"xxx | fdgcbe cefdb cefbgd gcb"}
		got := CountSimpleDigits(signalPattern7)
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Count output 8", func(t *testing.T) {
		signalPattern8 := []string{
			"xxx | fdgacbe cefdb cefbgd gadcb"}
		got := CountSimpleDigits(signalPattern8)
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Count full sample output", func(t *testing.T) {
		got := CountSimpleDigits(sampleSignalPattern)
		want := 26

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Count full part 1", func(t *testing.T) {
		data := readFullData(t, "day8.dat")
		fmt.Println(CountSimpleDigits(data))
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
