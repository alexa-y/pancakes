package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "5\n-\n-+\n+-\n+++\n--+-"
	fmt.Printf("Using input: %v\n", input)
	parsed, err := ParseInput(input)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	for i, p := range parsed {
		fmt.Printf("Case #%v: %v\n", i+1, CalculateMinFlips(p))
	}
}

func ParseInput(input string) ([][]int, error) {
	lines := strings.Split(input, "\n")
	stackCount, err := strconv.ParseInt(lines[0], 10, 32)

	if err != nil {
		return nil, errors.New("expected first line to be an integer")
	}

	stacks := make([][]int, stackCount)

	for i := 0; i < int(stackCount); i++ {
		stacks[i] = make([]int, len(lines[i+1]))
		for j, c := range lines[i+1] {
			if c == '+' {
				stacks[i][j] = 1
			}
		}
	}
	return stacks, nil
}

func CalculateMinFlips(stack []int) int {
	flips := 0
	for ; !isFlippedCorrectly(stack); flips++ {
		lowest := lowestBlank(stack)
		if stack[0] == 0 { // top is blank
			copy(stack[0:lowest+1], flip(stack[0:lowest+1]))
		} else {
			stack[0] = 0
		}
	}
	return flips
}

func lowestBlank(stack []int) int {
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] == 0 {
			return i
		}
	}
	return 0
}

func flip(stack []int) []int {
	reversed := make([]int, len(stack))
	for i, c := range stack {
		reversed[len(stack)-i-1] = ^c & 1
	}
	return reversed
}

func isFlippedCorrectly(stack []int) bool {
	for _, c := range stack {
		if c == 0 {
			return false
		}
	}
	return true
}
