package main

import (
	year2025 "aoc/2025"
	"fmt"
	"time"
)

type Solver interface {
	Solve([]string) string
}

type Date struct {
	year int
	day  int
	part int
}

func main() {
	tasks := map[Date]Solver{
		{year: 2025, day: 1, part: 1}: year2025.Day1Part1{},
		{year: 2025, day: 1, part: 2}: year2025.Day1Part2{},
		{year: 2025, day: 2, part: 1}: year2025.Day2Part1{},
		{year: 2025, day: 2, part: 2}: year2025.Day2Part2{},
	}
	var result string

	inputsFolderPath, err := getInputsFolderPath()
	if err != nil {
		panic(err)
	}

	for date, callable := range tasks {
		filePaths := getInputFilePaths(date, inputsFolderPath)

		for _, filepath := range filePaths {
			input := getDataFromFile(filepath)
			start := time.Now()
			result = callable.Solve(input)
			end := time.Now()
			fmt.Printf("Year %v Day %v Part %v\n", date.year, date.day, date.part)
			fmt.Printf("\tFile: %v\n", filepath)
			fmt.Printf("\tResult: %s\n", result)
			fmt.Printf("\tTook: %v\n", end.Sub(start))
		}
	}
}
