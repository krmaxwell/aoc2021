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

	singleLine := "edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc"

	wantA := [10]string{}
	wantA[0] = "abcdfg" // word will have six letters but is neither 6 nor 9
	wantA[1] = "cg"
	wantA[2] = "abcde"
	wantA[3] = "bcdeg"
	wantA[4] = "cefg"
	wantA[5] = "bdefg"  // word will have 5 letters and whatever fourdiff was
	wantA[6] = "abdefg" // word will have 6 letters and fourdiff but not all of 4
	wantA[7] = "bcg"
	wantA[8] = "abcdefg" // by definition
	wantA[9] = "bcdefg"

	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("Identify number %d in single line", i), func(t *testing.T) {
			got := DecodeSignal(singleLine)
			if got[i] != wantA[i] {
				t.Errorf("got %v want %v", got[i], wantA[i])
			}
		})
	}
	sampleValues := []int{
		8394,
		9781,
		1197,
		9361,
		4873,
		8418,
		4548,
		1625,
		8717,
		4315,
	}
	for i := 0; i < len(sampleValues); i++ {

		t.Run(fmt.Sprintf("Calculate output value for sample line %d", i), func(t *testing.T) {
			key := DecodeSignal(sampleSignalPattern[i])
			got := OutputValue(sampleSignalPattern[i], key)
			if got != sampleValues[i] {
				t.Errorf("got %d want %d", got, sampleValues[i])
			}
		})
	}

	t.Run("Calculate fourdiff in single line", func(t *testing.T) {
		got := wordsDiff("cefg", "cg")
		want := "ef"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Calculate all output values for sample input", func(t *testing.T) {
		got := 0
		for _, s := range sampleSignalPattern {
			key := DecodeSignal(s)
			value := OutputValue(s, key)
			got += value
		}

		want := 61229

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Count full part 2", func(t *testing.T) {
		data := readFullData(t, "day8.dat")
		got := 0
		for _, s := range data {
			key := DecodeSignal(s)
			value := OutputValue(s, key)
			got += value
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
