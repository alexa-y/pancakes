package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	stack         []int
	expectedFlips int
}

func compareIntSlices(s1, s2 []int) int {
	if len(s1) > len(s2) {
		return 1
	} else if len(s2) > len(s1) {
		return -1
	}

	for i := range s1 {
		if s1[i] > s2[i] {
			return 1
		} else if s2[i] > s1[i] {
			return -1
		}
	}

	return 0
}

func TestParseInput(t *testing.T) {
	input := "5\n-\n-+\n+-\n+++\n--+-"
	expectedParsed := [][]int{{0}, {0, 1}, {1, 0}, {1, 1, 1}, {0, 0, 1, 0}}

	parsed, err := ParseInput(input)

	if err != nil {
		t.Errorf("ParseInput error: %v", err)
	}

	for i := range expectedParsed {
		t.Run(fmt.Sprintf("Sample %v", i), func(t *testing.T) {
			if compareIntSlices(parsed[i], expectedParsed[i]) != 0 {
				t.Errorf("Expected %v to be %v", parsed, expectedParsed)
			}
		})
	}
}

func TestParseInputBadInput(t *testing.T) {
	input := "abc\n-"
	_, err := ParseInput(input)

	if err == nil {
		t.Errorf("Expected %v to produce an error", input)
	}
}

func TestCalculateMinFlips(t *testing.T) {
	tests := []TestCase{
		{stack: []int{0}, expectedFlips: 1},
		{stack: []int{0, 1}, expectedFlips: 1},
		{stack: []int{1, 0}, expectedFlips: 2},
		{stack: []int{1, 1, 1}, expectedFlips: 0},
		{stack: []int{0, 0, 1, 0}, expectedFlips: 3},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Testing stack %v", test.stack), func(t *testing.T) {
			flips := CalculateMinFlips(test.stack)
			if flips != test.expectedFlips {
				t.Errorf("Expected %v flips, got %v flips", test.expectedFlips, flips)
			}
		})
	}
}
