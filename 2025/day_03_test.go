package year2025

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestDay03Part01Solve(t *testing.T) {

// }

func TestGetMaxNumberFromLine(t *testing.T) {
	assert.Equal(t, 98, getMaxNumberFromLine("987654321111111"))
	assert.Equal(t, 89, getMaxNumberFromLine("811111111111119"))
	assert.Equal(t, 78, getMaxNumberFromLine("234234234234278"))
	assert.Equal(t, 92, getMaxNumberFromLine("818181911112111"))
}
