package year2024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestDay02Part01Solve(t *testing.T) {
// 	assert.True(t, solve([]int{7, 6, 4, 2, 1}))
// 	assert.False(t, solve([]int{1, 2, 7, 8, 9}))
// 	assert.False(t, solve([]int{9, 7, 6, 2, 1}))
// 	assert.False(t, solve([]int{1, 3, 2, 4, 5}))
// 	assert.False(t, solve([]int{8, 6, 4, 4, 1}))
// 	assert.True(t, solve([]int{1, 3, 6, 7, 9}))
// }

func TestReformatInputDay2(t *testing.T) {
	result, err := reformatInputDay2("7 6 4 2 1")
	assert.Equal(t, []int{7, 6, 4, 2, 1}, result)
	assert.Nil(t, err)

	result, err = reformatInputDay2("777777 6 4 213 1")
	assert.Equal(t, []int{777777, 6, 4, 213, 1}, result)
	assert.Nil(t, err)

	result, err = reformatInputDay2("anc6 4 213 1")
	assert.NotNil(t, err)
}

func TestGetCorruptedInd(t *testing.T) {
	result, hasCorruptedInd := getCorruptedIds([]int{7, 6, 4, 2, 1})
	assert.False(t, hasCorruptedInd)
	assert.Equal(t, []int{}, result)

	result, hasCorruptedInd = getCorruptedIds([]int{1, 2, 7, 8, 9})
	assert.True(t, hasCorruptedInd)
	assert.Equal(t, []int{1, 2}, result)

	result, hasCorruptedInd = getCorruptedIds([]int{19, 7, 6, 2, 1})
	assert.True(t, hasCorruptedInd)
	assert.Equal(t, []int{0, 1}, result)

	result, hasCorruptedInd = getCorruptedIds([]int{19, 7, 18, 16, 15})
	assert.True(t, hasCorruptedInd)
	assert.Equal(t, []int{0, 1}, result)

	result, hasCorruptedInd = getCorruptedIds([]int{1, 3, 2, 4, 5})
	assert.True(t, hasCorruptedInd)
	assert.Equal(t, []int{0, 1, 2}, result)

	result, hasCorruptedInd = getCorruptedIds([]int{8, 6, 4, 4, 1})
	assert.True(t, hasCorruptedInd)
	assert.Equal(t, []int{2, 3}, result)

	result, hasCorruptedInd = getCorruptedIds([]int{1, 3, 6, 7, 9})
	assert.False(t, hasCorruptedInd)
	assert.Equal(t, []int{}, result)
}

func TestRemoveElement(t *testing.T) {
	result := removeElement([]int{7, 6, 4, 2, 1}, 1)
	assert.Equal(t, []int{7, 4, 2, 1}, result)

	result = removeElement([]int{7, 6, 4, 2, 1}, 0)
	assert.Equal(t, []int{6, 4, 2, 1}, result)

	result = removeElement([]int{7, 6, 4, 2, 1}, 4)
	assert.Equal(t, []int{7, 6, 4, 2}, result)
}

func TestIsDampenedReportSafe(t *testing.T) {
	result := isDampenedReportSafe([]int{7, 6, 4, 2, 1})
	assert.True(t, result)

	result = isDampenedReportSafe([]int{1, 2, 7, 8, 9})
	assert.False(t, result)

	result = isDampenedReportSafe([]int{19, 7, 6, 2, 1})
	assert.False(t, result)

	result = isDampenedReportSafe([]int{19, 7, 6, 4, 3})
	assert.True(t, result)

	result = isDampenedReportSafe([]int{19, 7, 18, 16, 15})
	assert.True(t, result)

	result = isDampenedReportSafe([]int{1, 3, 2, 4, 5})
	assert.True(t, result)

	result = isDampenedReportSafe([]int{8, 6, 4, 4, 1})
	assert.True(t, result)

	result = isDampenedReportSafe([]int{8, 6, 4, 4, 4})
	assert.False(t, result)

	result = isDampenedReportSafe([]int{1, 3, 6, 7, 9})
	assert.True(t, result)
}
