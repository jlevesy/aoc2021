package day3

import (
	"errors"
	"math"
)

type PowerReport struct {
	GammaRate, EpsilonRate int
}

func EvaluatePowerConsumption(input [][]uint8) PowerReport {
	oneCounts := countOnes(input)

	gammaRateRaw := findCommonBitPerRow(oneCounts, len(input))
	espilonRateRaw := bitNot(gammaRateRaw)

	return PowerReport{
		GammaRate:   toDecimal(gammaRateRaw),
		EpsilonRate: toDecimal(espilonRateRaw),
	}
}

type LifeSupportReport struct {
	OxygenRating, CO2Rating int
}

var (
	errCannotDetermineRow = errors.New("cannot determine row")
)

func EvaluateLifeSupport(input [][]uint8) (LifeSupportReport, error) {
	oxygenRatingRow, err := findInterestingRow(input, mostCommonBit)

	if err != nil {
		return LifeSupportReport{}, err
	}

	co2ratingRow, err := findInterestingRow(input, mostUncommonBit)

	if err != nil {
		return LifeSupportReport{}, err
	}

	return LifeSupportReport{
		OxygenRating: toDecimal(oxygenRatingRow),
		CO2Rating:    toDecimal(co2ratingRow),
	}, nil
}

type bitCriteria func([][]uint8, int) uint8

func mostCommonBit(input [][]uint8, column int) uint8 {
	oneCounts := countOnes(input)

	// Force float64 division and comparison to control rounding.
	if float64(oneCounts[column]) >= math.Round(float64(len(input))/2) {
		return 1
	}

	return 0
}

func mostUncommonBit(input [][]uint8, column int) uint8 {
	if mostCommonBit(input, column) == 1 {
		return 0
	}

	return 1
}

func findInterestingRow(input [][]uint8, crit bitCriteria) ([]uint8, error) {
	if len(input) == 0 {
		return nil, errCannotDetermineRow
	}

	for i := 0; i < len(input[0]); i++ {
		if len(input) == 1 {
			break
		}

		input = filterByValue(input, i, crit(input, i))
	}

	if len(input) != 1 {
		return nil, errCannotDetermineRow
	}

	return input[0], nil
}

func filterByValue(input [][]uint8, column int, value uint8) [][]uint8 {
	var result [][]uint8

	for _, row := range input {
		if row[column] == value {
			result = append(result, row)
		}
	}

	return result
}

func countOnes(input [][]uint8) []int {
	var count []int

	for _, row := range input {
		// initialize the count slice only when we know the length of a row.
		if count == nil {
			count = make([]int, len(row))
		}

		for i, b := range row {
			if b > 0 {
				count[i]++
			}
		}
	}

	return count
}

func findCommonBitPerRow(counts []int, totalRows int) []uint8 {
	result := make([]uint8, len(counts))

	for i, count := range counts {
		// If the count of ones is supperior than the
		// half of the total rows, it means that it is the most common.
		// Otherwise let the slice at its zero value.
		if count > totalRows/2 {
			result[i] = 1
		}
	}

	return result
}

func bitNot(in []uint8) []uint8 {
	result := make([]uint8, len(in))

	for i, v := range in {
		// Only need to set positive bits.
		if v == 0 {
			result[i] = 1
		}
	}

	return result
}

func toDecimal(in []uint8) int {
	var result float64

	for i, v := range in {
		if v == 0 {
			continue
		}

		result += math.Pow(2, float64(len(in)-i-1))

	}

	return int(math.Round(result))
}
