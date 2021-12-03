package day2_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/jlevesy/aoc/day2"
	"github.com/jlevesy/aoc/pkg/input"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleCommands = []day2.Command{
	{
		Action: day2.ActionForward,
		Amount: 5,
	},
	{
		Action: day2.ActionDown,
		Amount: 5,
	},
	{
		Action: day2.ActionForward,
		Amount: 8,
	},
	{
		Action: day2.ActionUp,
		Amount: 3,
	},
	{
		Action: day2.ActionDown,
		Amount: 8,
	},
	{
		Action: day2.ActionForward,
		Amount: 2,
	},
}

func TestEvaluatePosition(t *testing.T) {
	pos := day2.EvaluatePosition(exampleCommands)

	assert.Equal(t, 150, pos.Depth*pos.Horizontal)
}

func TestEvaluatePosition_Answer(t *testing.T) {
	pos := day2.EvaluatePosition(readInput(t))

	t.Logf(
		"Horizontal is %d, Depth is %d, multiplication is %d",
		pos.Horizontal,
		pos.Depth,
		pos.Horizontal*pos.Depth,
	)
}

func TestEvaluatePositionWithAim(t *testing.T) {
	pos := day2.EvaluatePositionWithAim(exampleCommands)

	assert.Equal(t, 900, pos.Depth*pos.Horizontal)
}

func TestEvaluatePositionWithAim_Answer(t *testing.T) {
	pos := day2.EvaluatePositionWithAim(readInput(t))

	t.Logf(
		"Horizontal is %d, Depth is %d, multiplication is %d",
		pos.Horizontal,
		pos.Depth,
		pos.Horizontal*pos.Depth,
	)
}

func readInput(t *testing.T) []day2.Command {
	t.Helper()

	var result []day2.Command

	err := input.ReadInput("./fixtures/input.txt", func(line string) error {
		splits := strings.Split(line, " ")

		if len(splits) != 2 {
			return fmt.Errorf("unexpected amount of splits, expected 2 got %d", len(splits))
		}

		amount, err := strconv.Atoi(splits[1])
		if err != nil {
			return err
		}

		action, err := parseAction(splits[0])
		if err != nil {
			return err
		}

		result = append(result, day2.Command{Action: action, Amount: amount})

		return nil
	})
	require.NoError(t, err)

	return result
}

func parseAction(action string) (day2.Action, error) {
	switch action {
	case "forward":
		return day2.ActionForward, nil
	case "down":
		return day2.ActionDown, nil
	case "up":
		return day2.ActionUp, nil
	default:
		return day2.ActionUnknown, fmt.Errorf("unsupported action %q", action)
	}

}
