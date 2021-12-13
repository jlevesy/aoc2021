package day4

import (
	"errors"
)

type Board struct {
	// values represent the value for a board
	values [][]int64
	marked map[int64]struct{}
}

func NewBoard(values [][]int64) Board {
	return Board{
		values: values,
		marked: map[int64]struct{}{},
	}
}

func (b *Board) Push(row []int64) {
	b.values = append(b.values, row)
}

// detects if the value is being used at least once in the board.
// and stores its value into the marked map.
func (b *Board) mark(in int64) {
	for _, row := range b.values {
		for _, v := range row {
			if v == in {
				b.marked[v] = struct{}{}
				return
			}
		}
	}
}

// returns true if the board has a complete row or column according to
// its marked map.
func (b *Board) wins() bool {
	var (
		rowMarkCounts    = make([]int, len(b.values))
		columnMarkCounts = make([]int, len(b.values))
	)

	// Count per row and per column how many marks match.
	for rowIdx, row := range b.values {
		for colIdx, val := range row {
			if _, ok := b.marked[val]; ok {
				rowMarkCounts[rowIdx]++
				columnMarkCounts[colIdx]++
			}
		}
	}

	// If one of the count per row equals the width of the grid
	// then it means that we have a complete row.
	for _, count := range rowMarkCounts {
		if count == len(b.values) {
			return true
		}
	}

	// If one of the count per column equals the width of the grid
	// then it means that we have a complete column.
	for _, count := range columnMarkCounts {
		if count == len(b.values) {
			return true
		}
	}

	// Otherwise, we did not win just yet.
	return false
}

func (b *Board) sumUnmarked() int64 {
	var sum int64

	for _, row := range b.values {
		for _, v := range row {
			if _, ok := b.marked[v]; !ok {
				sum += v
			}
		}
	}

	return sum
}

func PlayBingo(boards []Board, draws []int64) (int64, error) {
	for _, draw := range draws {
		for _, board := range boards {
			board.mark(draw)

			if board.wins() {
				return draw * board.sumUnmarked(), nil
			}
		}
	}

	return 0, errors.New("no winner after all draws")
}

func WinsLast(boards []Board, draws []int64) (int64, error) {
	var (
		winningCount int
		wonBoards    = make([]bool, len(boards))
	)

	for _, draw := range draws {
		for i, board := range boards {
			if wonBoards[i] {
				continue
			}

			board.mark(draw)

			if board.wins() {
				winningCount += 1
				wonBoards[i] = true
			}

			if len(boards) == winningCount {
				return draw * board.sumUnmarked(), nil
			}
		}
	}

	return 0, errors.New("no winner after all draws")
}
