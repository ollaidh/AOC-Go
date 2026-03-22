package year2025

import (
	"strconv"
	"strings"
)

type Day2Part1 struct{}

func (day Day2Part1) Solve(input []string) string {
	ranges := parseInputDayTwo(input[0])
	result := Solve(ranges, isIDValidPartOne)
	return strconv.Itoa(result)
}

type Day2Part2 struct{}

func (day Day2Part2) Solve(input []string) string {
	ranges := parseInputDayTwo(input[0])
	result := Solve(ranges, isIDValidPartTwo)
	return strconv.Itoa(result)
}

type Range struct {
	minValue int
	maxValue int
}

func parseInputDayTwo(input string) []Range {
	result := []Range{}
	inputRanges := strings.Split(input, ",")
	for _, currRange := range inputRanges {
		currCurrRange := strings.Split(currRange, "-")
		start, _ := strconv.Atoi(currCurrRange[0])
		end, _ := strconv.Atoi(currCurrRange[1])
		result = append(result, Range{minValue: start, maxValue: end})
	}
	return result
}

func getDigitsNumber(number int) int {
	counter := 0
	for number > 0 {
		number /= 10
		counter += 1
	}
	return counter
}

func powInt(number int, power int) int {
	result := number
	for range power - 1 {
		result *= number
	}
	return result
}

func isIDValidPartOne(id int) bool {
	digitsNumber := getDigitsNumber(id)
	if digitsNumber%2 != 0 {
		return true
	}
	halfLength := powInt(10, digitsNumber/2)
	firstHalf := id / halfLength
	secondHalf := id % halfLength

	return firstHalf != secondHalf

}

func isIDValidPartTwo(id int) bool {
	word := strconv.Itoa(id)
	length := len(word)
	for i := 1; i <= length/2; i++ {
		if length%i != 0 {
			continue
		}
		perfectWord := strings.Repeat(word[0:i], length/i)
		if perfectWord == word {
			return false
		}
	}
	return true
}

func Solve(ranges []Range, check func(int) bool) int {
	result := 0
	for i := range ranges {
		for id := ranges[i].minValue; id <= ranges[i].maxValue; id++ {
			if !check(id) {
				result += id
			}
		}
	}
	return result
}
