package day4_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/jlevesy/aoc/day4"
	"github.com/jlevesy/aoc/pkg/input"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleBoards = []day4.Board{
	day4.NewBoard(
		[][]int64{
			{22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7},
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		},
	),
	day4.NewBoard(
		[][]int64{
			{3, 15, 0, 2, 22},
			{9, 18, 13, 17, 5},
			{19, 8, 7, 25, 23},
			{20, 11, 10, 24, 4},
			{14, 21, 16, 12, 6},
		},
	),
	day4.NewBoard(
		[][]int64{
			{14, 21, 17, 24, 4},
			{10, 16, 15, 9, 19},
			{18, 8, 23, 26, 20},
			{22, 11, 13, 6, 5},
			{2, 0, 12, 3, 7},
		},
	),
}

var exampleDraws = []int64{
	7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1,
}

func TestPlayBingo(t *testing.T) {
	result, err := day4.PlayBingo(exampleBoards, exampleDraws)
	require.NoError(t, err)
	assert.Equal(t, int64(4512), result)
}

func TestPlayBingo_Answer(t *testing.T) {
	boards, draws := readInput(t)

	result, err := day4.PlayBingo(boards, draws)
	require.NoError(t, err)

	t.Log("Result is", result)
}

func TestWinsLast(t *testing.T) {
	result, err := day4.WinsLast(exampleBoards, exampleDraws)
	require.NoError(t, err)
	assert.Equal(t, int64(1924), result)
}

func TestWinsLast_Answer(t *testing.T) {
	boards, draws := readInput(t)

	result, err := day4.WinsLast(boards, draws)
	require.NoError(t, err)

	t.Log("Result is", result)
}

func readInput(t *testing.T) ([]day4.Board, []int64) {
	t.Helper()

	var (
		err    error
		draws  []int64
		boards []day4.Board
	)

	err = input.ReadInput("./fixtures/input.txt", func(line string) error {
		if len(draws) == 0 {
			draws, err = parseLine(strings.Split(line, ","))
			return err
		}

		if line == "" {
			boards = append(boards, day4.NewBoard(nil))
			return nil
		}

		row, err := parseLine(strings.Split(line, " "))

		if err != nil {
			return err
		}

		boards[len(boards)-1].Push(row)

		return nil
	})
	require.NoError(t, err)

	return boards, draws
}

func parseLine(in []string) ([]int64, error) {
	var line []int64

	for _, s := range in {
		if s == "" {
			continue
		}

		v, err := strconv.ParseInt(s, 10, 64)

		if err != nil {
			return nil, err
		}

		line = append(line, v)
	}

	return line, nil
}
