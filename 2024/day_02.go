package year2024

import (
	"aoc/internal/mathx"
	"strconv"
	"strings"
)

type Day2Part1 struct{}

func reformatInputDay2(input string) ([]int, error) {
	parts := strings.Fields(input)
	result := make([]int, len(parts))
	for i, part := range parts {
		v, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		result[i] = v
	}
	return result, nil
}

func isReportSafe(input []int) bool {
	sign := 0 // 0 - Unknown, 1 - Positive, -1 - Negative
	for i := range len(input) - 1 {
		diff := input[i] - input[i+1]
		diffAbs := mathx.AbsInt(diff)
		if diffAbs == 0 || diffAbs > 3 {
			return false
		}
		currSign := diff / diffAbs
		if sign != 0 && currSign != sign {
			return false
		}
		sign = currSign
	}
	return true
}

func (day Day2Part1) Solve(input []string) string {
	safeReportsCount := 0
	for _, line := range input {
		report, _ := reformatInputDay2(line)
		if isReportSafe(report) {
			safeReportsCount += 1
		}
	}
	return strconv.Itoa(safeReportsCount)
}
