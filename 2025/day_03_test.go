package year2025

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestDay03Part01Solve(t *testing.T) {

// }

func TestGetMaxNumberFromLine(t *testing.T) {
	assert.Equal(t, 98, getMaxNumberFromLine("987654321111111", 2))
	assert.Equal(t, 89, getMaxNumberFromLine("811111111111119", 2))
	assert.Equal(t, 78, getMaxNumberFromLine("234234234234278", 2))
	assert.Equal(t, 92, getMaxNumberFromLine("818181911112111", 2))

	assert.Equal(t, 987654321111, getMaxNumberFromLine("987654321111111", 12))
	assert.Equal(t, 811111111119, getMaxNumberFromLine("811111111111119", 12))
	assert.Equal(t, 434234234278, getMaxNumberFromLine("234234234234278", 12))
	assert.Equal(t, 888911112111, getMaxNumberFromLine("818181911112111", 12))
}

func TestGetMaxDigit(t *testing.T) {
	var idx int
	var maxVal int

	idx, maxVal = getMaxDigit(0, 14, "987654321111111")
	assert.Equal(t, 0, idx)
	assert.Equal(t, 9, maxVal)

	idx, maxVal = getMaxDigit(1, 14, "987654321111111")
	assert.Equal(t, 1, idx)
	assert.Equal(t, 8, maxVal)

	idx, maxVal = getMaxDigit(0, 14, "811111111111119")
	assert.Equal(t, 14, idx)
	assert.Equal(t, 9, maxVal)

	idx, maxVal = getMaxDigit(14, 14, "811111111111119")
	assert.Equal(t, 14, idx)
	assert.Equal(t, 9, maxVal)

	idx, maxVal = getMaxDigit(0, 14, "234234234234278")
	assert.Equal(t, 14, idx)
	assert.Equal(t, 8, maxVal)

	idx, maxVal = getMaxDigit(5, 14, "818181911112111")
	assert.Equal(t, 6, idx)
	assert.Equal(t, 9, maxVal)
}
