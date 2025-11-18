package main

import (
	_ "embed"
	"testing"
)

func TestPart1(t *testing.T) {
	input := ""
	result := Part1(input)
	expected := 1
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2(t *testing.T) {
	input := ""
	result := Part2(input)
	expected := 2
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
