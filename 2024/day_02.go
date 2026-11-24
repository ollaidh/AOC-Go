package year2024

import (
	"aoc/internal/mathx"
	"slices"
	"strconv"
	"strings"
)

type Day2Part1 struct{}
type Day2Part2 struct{}

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

func getCorruptedIds(input []int) ([]int, bool) {
	sign := 0 // 0 - Unknown, 1 - Positive, -1 - Negative
	for i := range len(input) - 1 {
		diff := input[i] - input[i+1]
		diffAbs := mathx.AbsInt(diff)
		if diffAbs == 0 || diffAbs > 3 {
			return []int{i, i + 1}, true
		}
		currSign := diff / diffAbs
		if sign != 0 && currSign != sign {
			return []int{i - 1, i, i + 1}, true
		}
		sign = currSign
	}
	return []int{}, false
}

func removeElement(line []int, idx int) []int {
	return slices.Concat(line[:idx], line[idx+1:])
}

func isDampenedReportSafe(report []int) bool {
	idxs, hasCorruptedIdx := getCorruptedIds(report)
	if !hasCorruptedIdx {
		return true
	}

	for _, idx := range idxs {
		reportWithoutCorruptedIdx := removeElement(report, idx)
		_, hasCorrupted := getCorruptedIds(reportWithoutCorruptedIdx)
		if !hasCorrupted {
			return true
		}
	}

	return false

}

func (day Day2Part1) Solve(input []string) string {
	safeReportsCount := 0
	for _, line := range input {
		report, _ := reformatInputDay2(line)
		_, hasCorruptedIdx := getCorruptedIds(report)
		if !hasCorruptedIdx {
			safeReportsCount += 1
		}
	}
	return strconv.Itoa(safeReportsCount)
}

func (day Day2Part2) Solve(input []string) string {
	safeReportsCount := 0
	for _, line := range input {
		report, _ := reformatInputDay2(line)
		if isDampenedReportSafe(report) {
			safeReportsCount += 1
		}

	}
	return strconv.Itoa(safeReportsCount)
}
