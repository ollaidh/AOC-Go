package year2024

import (
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

func (day Day2Part1) Solve(input []string) string {
	return ""
}
