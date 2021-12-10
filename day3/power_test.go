package day3_test

import (
	"fmt"
	"testing"

	"github.com/jlevesy/aoc/day3"
	"github.com/jlevesy/aoc/pkg/input"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var example = [][]uint8{
	{0, 0, 1, 0, 0},
	{1, 1, 1, 1, 0},
	{1, 0, 1, 1, 0},
	{1, 0, 1, 1, 1},
	{1, 0, 1, 0, 1},
	{0, 1, 1, 1, 1},
	{0, 0, 1, 1, 1},
	{1, 1, 1, 0, 0},
	{1, 0, 0, 0, 0},
	{1, 1, 0, 0, 1},
	{0, 0, 0, 1, 0},
	{0, 1, 0, 1, 0},
}

func TestEvaluatePowerConsumption(t *testing.T) {
	report := day3.EvaluatePowerConsumption(example)
	assert.Equal(t, 198, report.GammaRate*report.EpsilonRate)
}

func TestEvaluatePowerConsumption_Answer(t *testing.T) {
	input := readInput(t)
	report := day3.EvaluatePowerConsumption(input)

	t.Logf(
		"GammaRate %d, EpsilonRate %d, multiplication is %d",
		report.GammaRate,
		report.EpsilonRate,
		report.GammaRate*report.EpsilonRate,
	)
}

func TestEvaluateLifeSupport(t *testing.T) {
	report, err := day3.EvaluateLifeSupport(example)
	require.NoError(t, err)
	assert.Equal(t, 23, report.OxygenRating)
	assert.Equal(t, 10, report.CO2Rating)
}

func TestEvaluateLifeSupport_Answer(t *testing.T) {
	input := readInput(t)
	report, err := day3.EvaluateLifeSupport(input)
	require.NoError(t, err)

	t.Logf(
		"OxygenRating %d, CO2Rating %d, multiplication is %d",
		report.OxygenRating,
		report.CO2Rating,
		report.OxygenRating*report.CO2Rating,
	)
}

func readInput(t *testing.T) [][]uint8 {
	t.Helper()

	var result [][]uint8

	err := input.ReadInput("./fixtures/input.txt", func(line string) error {
		bitLine := make([]uint8, len(line))

		for i, v := range line {
			switch v {
			case '1':
				bitLine[i] = 1
			case '0':
			default:
				return fmt.Errorf("unexpected value %q at index %q", v, i)
			}
		}

		result = append(result, bitLine)

		return nil
	})
	require.NoError(t, err)

	return result

}
