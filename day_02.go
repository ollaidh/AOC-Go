package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Day2Part1 struct{}

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

func readInputDayTwo(inputPath string) string {
	f, _ := os.Open(inputPath)
	scanner := bufio.NewScanner(f)
	result := scanner.Text()
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

func isIDValid(id int) bool {
	digitsNumber := getDigitsNumber(id)
	if digitsNumber%2 != 0 {
		return true
	}
	halfLength := powInt(10, digitsNumber/2)
	firstHalf := id / halfLength
	secondHalf := id % halfLength

	return firstHalf != secondHalf

}

func solve(ranges []Range) int {
	result := 0
	for i := range ranges {
		for id := ranges[i].minValue; id <= ranges[i].maxValue; id++ {
			if !isIDValid(id) {
				result += id
			}
		}
	}
	return result
}
