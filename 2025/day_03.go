package year2025

import "strconv"

type Day3Part1 struct{}

func (day Day3Part1) Solve(input []string) string {
	sum := 0
	for _, line := range input {
		sum += getMaxNumberFromLine(line, 2)
	}
	return strconv.Itoa(sum)
}

type Day3Part2 struct{}

func (day Day3Part2) Solve(input []string) string {
	sum := 0
	for _, line := range input {
		sum += getMaxNumberFromLine(line, 12)
	}
	return strconv.Itoa(sum)
}

func getMaxDigit(start int, end int, line string) (int, int) {
	var maxLetter byte
	maxLetterIdx := 0

	for i := start; i <= end; i++ {
		if line[i] > maxLetter {
			maxLetter = line[i]
			maxLetterIdx = i
		}
	}
	return maxLetterIdx, int(maxLetter - '0')
}

func getMaxNumberFromLine(line string, digitCount int) int {
	result := 0

	start := 0
	end := len(line) - digitCount
	for i := 0; i < digitCount; i++ {
		idx, val := getMaxDigit(start, end, line)
		result = result*10 + val
		start = idx + 1
		end += 1
	}

	return result
}
