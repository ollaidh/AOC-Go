package year2025

import "strconv"

type Day3Part1 struct{}

func (day Day3Part1) Solve(input []string) string {
	sum := 0
	for _, line := range input {
		sum += getMaxNumberFromLine(line)
	}
	return strconv.Itoa(sum)
}

func getMaxNumberFromLine(line string) int {
	var firstMaxLetter byte
	firstMaxLetterIdx := 0
	for i := 0; i < len(line)-1; i++ {
		if line[i] > firstMaxLetter {
			firstMaxLetter = line[i]
			firstMaxLetterIdx = i
		}
	}
	secondMaxLetter := line[firstMaxLetterIdx+1]
	for j := firstMaxLetterIdx + 1; j < len(line); j++ {
		if line[j] > secondMaxLetter {
			secondMaxLetter = line[j]
		}
	}
	return int(firstMaxLetter-'0')*10 + int(secondMaxLetter-'0')
}
