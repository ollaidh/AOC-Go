package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInputDayTwo(t *testing.T) {
	initialInput := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	result := parseInputDayTwo(initialInput)
	assert.Equal(t, 11, len(result))
	assert.Equal(t, 11, result[0].minValue)
	assert.Equal(t, 22, result[0].maxValue)
	assert.Equal(t, 95, result[1].minValue)
	assert.Equal(t, 115, result[1].maxValue)
	assert.Equal(t, 998, result[2].minValue)
	assert.Equal(t, 1012, result[2].maxValue)
	assert.Equal(t, 1188511880, result[3].minValue)
	assert.Equal(t, 1188511890, result[3].maxValue)
	assert.Equal(t, 222220, result[4].minValue)
	assert.Equal(t, 222224, result[4].maxValue)
	assert.Equal(t, 1698522, result[5].minValue)
	assert.Equal(t, 1698528, result[5].maxValue)
	assert.Equal(t, 446443, result[6].minValue)
	assert.Equal(t, 446449, result[6].maxValue)
	assert.Equal(t, 38593856, result[7].minValue)
	assert.Equal(t, 38593862, result[7].maxValue)
	assert.Equal(t, 565653, result[8].minValue)
	assert.Equal(t, 565659, result[8].maxValue)
	assert.Equal(t, 824824821, result[9].minValue)
	assert.Equal(t, 824824827, result[9].maxValue)
	assert.Equal(t, 2121212118, result[10].minValue)
	assert.Equal(t, 2121212124, result[10].maxValue)

}

func TestPowInt(t *testing.T) {
	assert.Equal(t, 1000, powInt(10, 3))
	assert.Equal(t, 10, powInt(10, 1))
	assert.Equal(t, 16, powInt(2, 4))
	assert.Equal(t, 1, powInt(1, 10))
}

func TestIsIDValidPartOne(t *testing.T) {
	assert.True(t, isIDValidPartOne(12345))
	assert.True(t, isIDValidPartOne(123456))
	assert.False(t, isIDValidPartOne(123123))
	assert.False(t, isIDValidPartOne(111111))
	assert.True(t, isIDValidPartOne(11111))
	assert.True(t, isIDValidPartOne(3))
	assert.False(t, isIDValidPartOne(22))
	assert.True(t, isIDValidPartOne(100))
	assert.False(t, isIDValidPartOne(100100))
}

func TestIsIDValidPartTwo(t *testing.T) {
	assert.True(t, isIDValidPartTwo(12345))
	assert.True(t, isIDValidPartTwo(123456))
	assert.False(t, isIDValidPartTwo(123123))
	assert.False(t, isIDValidPartTwo(123123123))
	assert.False(t, isIDValidPartTwo(1212121212))
	assert.False(t, isIDValidPartTwo(111111))
	assert.False(t, isIDValidPartTwo(11111))
	assert.True(t, isIDValidPartTwo(3))
	assert.False(t, isIDValidPartTwo(22))
	assert.True(t, isIDValidPartTwo(100))
	assert.False(t, isIDValidPartTwo(100100))
}

func TestSolvePartOne(t *testing.T) {
	initialInput := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	ranges := parseInputDayTwo(initialInput)
	assert.Equal(t, solve(ranges, isIDValidPartOne), 1227775554)
}
