package day1

func CountIncreases(input []int) int {
	return countPositives(deltas(input))
}

func CountIncreasesSlidingWindows(input []int) int {
	return countPositives(deltas(sumWindows(input)))
}

func countPositives(input []int) int {
	var count int

	for _, value := range input {
		if value > 0 {
			count++
		}
	}

	return count
}

func sumWindows(input []int) []int {
	sums := make([]int, len(input))

	for i := 1; i < len(input)-1; i++ {
		sums[i-1] = sum(input[i-1 : i+2])
	}

	return sums
}

func sum(in []int) int {
	var result int

	for _, v := range in {
		result += v
	}

	return result
}

func deltas(input []int) []int {
	deltas := make([]int, len(input)-1)

	for i := 1; i < len(input); i++ {
		deltas[i-1] = input[i] - input[i-1]
	}

	return deltas
}
