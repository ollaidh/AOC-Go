package year2024

import (
	"sort"
	"strconv"
	"strings"
)

type Day1Part1 struct{}

func (day Day1Part1) Solve(input []string) string {
	formattedInput, err := reformatInput(input)
	if err != nil {
		panic(err)
	}
	result := solvePart1(formattedInput)
	return strconv.Itoa(result)
}

type Day1Part2 struct{}

func (day Day1Part2) Solve(input []string) string {
	formattedInput, err := reformatInput(input)
	if err != nil {
		panic(err)
	}
	result := solvePart2(formattedInput)
	return strconv.Itoa(result)
}

type Day1Input struct {
	Lefties  []int
	Righties []int
}

func reformatInput(input []string) (Day1Input, error) {
	result := Day1Input{
		Lefties:  make([]int, 0, len(input)),
		Righties: make([]int, 0, len(input)),
	}
	for _, row := range input {
		parts := strings.Fields(row)
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			return Day1Input{}, err
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			return Day1Input{}, err
		}
		result.Lefties = append(result.Lefties, left)
		result.Righties = append(result.Righties, right)
	}

	return result, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solvePart1(input Day1Input) int {
	sort.Ints(input.Lefties)
	sort.Ints(input.Righties)

	result := 0

	for i := 0; i < len(input.Lefties); i++ {
		result += abs(input.Lefties[i] - input.Righties[i])
	}

	return result
}

func solvePart2(input Day1Input) int {
	freqs := make(map[int]int)
	for _, val := range input.Righties {
		if _, ok := freqs[val]; !ok {
			freqs[val] = 0
		}
		freqs[val] += 1
	}

	result := 0

	for _, val := range input.Lefties {
		if _, ok := freqs[val]; !ok {
			freqs[val] = 0
		}
		result += val * freqs[val]
	}

	return result
}
