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
	result := solve(formattedInput)
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

func solve(input Day1Input) int {
	sort.Ints(input.Lefties)
	sort.Ints(input.Righties)

	result := 0

	for i := 0; i < len(input.Lefties); i++ {
		result += abs(input.Lefties[i] - input.Righties[i])
	}

	return result
}
