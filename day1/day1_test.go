package day1_test

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"testing"

	"github.com/jlevesy/aoc/day1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var example = []int{
	199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	263,
}

func TestCountIncreases(t *testing.T) {
	gotIncreases := day1.CountIncreases(example)
	assert.Equal(t, 7, gotIncreases)
}

func TestCountIncreases_Answer(t *testing.T) {
	input, err := readInput(t)
	require.NoError(t, err)

	t.Log("Total of increases", day1.CountIncreases(input))
}

func TestCountIncreasesSlidingWindows(t *testing.T) {
	gotIncreases := day1.CountIncreasesSlidingWindows(example)
	assert.Equal(t, 5, gotIncreases)
}

func TestCountIncreasesSlidingWindows_Answer(t *testing.T) {
	input, err := readInput(t)
	require.NoError(t, err)

	t.Log("Total of increases with sliding windows", day1.CountIncreasesSlidingWindows(input))
}

func readInput(t *testing.T) ([]int, error) {
	t.Helper()

	file, err := os.Open("./fixtures/input.txt")
	require.NoError(t, err)

	defer file.Close()

	return parseInput(file)
}

func parseInput(input io.Reader) ([]int, error) {
	var (
		scanner = bufio.NewScanner(input)
		result  []int
	)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		result = append(result, value)
	}

	return result, scanner.Err()
}
