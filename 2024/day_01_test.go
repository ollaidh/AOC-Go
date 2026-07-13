package year2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01Part01Solve(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	result := solve(
		Day1Input{
			Lefties:  left,
			Righties: right,
		},
	)
	assert.Equal(t, 11, result)

}
