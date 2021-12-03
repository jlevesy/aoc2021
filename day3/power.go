package day3

import (
	"math"
)

type Report struct {
	GammaRate, EpsilonRate int
}

func EvaluatePowerConsumption(input [][]uint8) Report {
	oneCounts := countOnes(input)

	gammaRateRaw := findCommonBitPerRow(oneCounts, len(input))
	espilonRateRaw := bitNot(gammaRateRaw)

	return Report{
		GammaRate:   toDecimal(gammaRateRaw),
		EpsilonRate: toDecimal(espilonRateRaw),
	}
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
