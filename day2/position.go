package day2

type Action int

const (
	ActionUnknown Action = iota
	ActionForward
	ActionUp
	ActionDown
)

type Command struct {
	Action Action
	Amount int
}

type Position struct {
	Horizontal, Depth int
}

func EvaluatePosition(commands []Command) Position {
	var pos Position

	for _, command := range commands {
		switch command.Action {
		case ActionForward:
			pos.Horizontal += command.Amount
		case ActionDown:
			pos.Depth += command.Amount
		case ActionUp:
			pos.Depth -= command.Amount
		}
	}

	return pos
}

func EvaluatePositionWithAim(commands []Command) Position {
	var (
		pos Position
		aim int
	)

	for _, command := range commands {
		switch command.Action {
		case ActionForward:
			pos.Horizontal += command.Amount
			pos.Depth += aim * command.Amount
		case ActionDown:
			aim += command.Amount
		case ActionUp:
			aim -= command.Amount
		}
	}

	return pos
}
