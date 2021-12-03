package day1_test

import (
	"strconv"
	"testing"

	"github.com/jlevesy/aoc/day1"
	"github.com/jlevesy/aoc/pkg/input"

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
	input := readInput(t)

	t.Log("Total of increases", day1.CountIncreases(input))
}

func TestCountIncreasesSlidingWindows(t *testing.T) {
	gotIncreases := day1.CountIncreasesSlidingWindows(example)
	assert.Equal(t, 5, gotIncreases)
}

func TestCountIncreasesSlidingWindows_Answer(t *testing.T) {
	input := readInput(t)

	t.Log("Total of increases with sliding windows", day1.CountIncreasesSlidingWindows(input))
}

func readInput(t *testing.T) []int {
	t.Helper()

	var result []int

	err := input.ReadInput("./fixtures/input.txt", func(line string) error {
		value, err := strconv.Atoi(line)
		if err != nil {
			return err
		}

		result = append(result, value)

		return nil
	})
	require.NoError(t, err)

	return result
}
