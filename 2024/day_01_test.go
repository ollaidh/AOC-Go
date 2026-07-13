package year2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01Part01Solve(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	result := solvePart1(
		Day1Input{
			Lefties:  left,
			Righties: right,
		},
	)
	assert.Equal(t, 11, result)
}

func TestDay01Part02Solve(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	result := solvePart2(
		Day1Input{
			Lefties:  left,
			Righties: right,
		},
	)
	assert.Equal(t, 31, result)
}
