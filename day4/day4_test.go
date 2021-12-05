package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestWinningRoundByRow(t *testing.T) {
	draw := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	got := Board{Layout: BoardLayout{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 4, 18, 5},
		{1, 12, 20, 15, 19},
	}}.CalcResult(draw)

	want := Board{Round: 14, HasWon: true}

	if got.Round != want.Round || got.HasWon != want.HasWon {
		t.Errorf("got %d want %d", got.Round, want.Round)
	}
}

func TestWinningRoundByColumn(t *testing.T) {
	draw := []int{21, 6, 1, 9, 22, 13, 8, 0, 11, 17}
	got := Board{Layout: BoardLayout{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 4, 18, 5},
		{1, 12, 20, 15, 19},
	}}.CalcResult(draw)

	want := Board{Round: 7, HasWon: true}

	if got.Round != want.Round || got.HasWon != want.HasWon {
		t.Errorf("got %d want %d", got.Round, want.Round)
	}
}

func TestLosingBoard(t *testing.T) {
	draw := []int{0}
	got := Board{Layout: BoardLayout{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 4, 18, 5},
		{1, 12, 20, 15, 19},
	}}.CalcResult(draw)

	want := false

	if got.HasWon != want {
		t.Errorf("got %v want %v", got.HasWon, want)
	}

}

func TestBoardScore(t *testing.T) {
	draw := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	board := Board{Layout: BoardLayout{
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7},
	}}.CalcResult(draw)
	got := board.Score
	want := 4512

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestPartOne(t *testing.T) {
	draw := []int{83, 69, 34, 46, 30, 23, 19, 75, 22, 37, 89, 78, 32, 39, 11, 44, 95, 43, 26, 48, 84, 53, 94, 88, 18, 40, 62, 35, 27, 42, 15, 2, 91, 20, 4, 64, 99, 71, 54, 97, 52, 36, 28, 7, 74, 45, 70, 86, 98, 1, 61, 50, 68, 6, 77, 8, 57, 47, 51, 72, 65, 3, 49, 24, 79, 13, 17, 92, 41, 80, 63, 67, 82, 90, 55, 0, 10, 93, 38, 21, 59, 73, 33, 31, 9, 76, 5, 66, 16, 58, 85, 87, 12, 29, 25, 14, 96, 56, 60, 81}
	layouts := getPuzzleInput(t, "day4.dat")

	earliestWin := 999999
	earliestWinningBoard := 0
	earliestWinningScore := 0
	latestWin := 0
	latestWinningBoard := 0
	latestWinningScore := 0

	for i, layout := range layouts {
		board := Board{Layout: layout}
		board = board.CalcResult(draw)
		if board.Round < earliestWin {
			earliestWin = board.Round
			earliestWinningBoard = i
			earliestWinningScore = board.Score
		}
		if board.Round > latestWin {
			latestWin = board.Round
			latestWinningBoard = i
			latestWinningScore = board.Score
		}
	}

	fmt.Printf("Board %d wins on round %d with score %d\n", earliestWinningBoard, earliestWin, earliestWinningScore)
	fmt.Printf("Board %d wins on round %d with score %d\n", latestWinningBoard, latestWin, latestWinningScore)
}

func getPuzzleInput(t *testing.T, fname string) []BoardLayout {
	t.Helper()
	f, err := os.Open(fname)
	if err != nil {
		t.Fatalf("Could not open %q: %q", fname, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	input := []string{}

	// read in the whole file
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	puzzle := []BoardLayout{}
	i := 0
	for i < len(input)-5 {
		currentBoard := BoardLayout{}
		for j := 0; j < 5; j++ {
			row := [5]int{}
			l := strings.Fields(input[i+j])
			for k := 0; k < 5; k++ {
				row[k], err = ParseIntBase(l[k], 10)
				if err != nil {
					t.Fatalf("Couldn't parse int: %q (%q)", err, l)
				}
			}
			currentBoard[j] = row
			i += 1
		}
		i += 1 // skip a line between
		puzzle = append(puzzle, currentBoard)
	}

	return puzzle
}

func ParseIntBase(s string, base int) (int, error) {
	i64, err := strconv.ParseInt(s, base, 0)
	return int(i64), err
}
