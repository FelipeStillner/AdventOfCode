package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

const (
	NOTHING   = 0
	SPLITTLER = -1
)

func Parse(input string) [][]int {
	var matrix [][]int

	lines := strings.Split(input, "\n")
	matrix = make([][]int, len(lines))
	for i := range matrix {
		matrix[i] = make([]int, len(lines[0]))
	}

	for i, line := range lines {
		for j, v := range line {
			switch v {
			case 'S':
				matrix[i][j] = 1
			case '^':
				matrix[i][j] = SPLITTLER
			default:
				matrix[i][j] = NOTHING
			}
		}
	}

	return matrix
}

func Part1(input string) int {
	matrix := Parse(input)
	sum := 0
	for i, line := range matrix {
		for j, val := range line {
			if val <= 0 || i == len(matrix)-1 {
				continue
			}

			switch matrix[i+1][j] {
			case NOTHING:
				matrix[i+1][j] = 1
			case SPLITTLER:
				sum++
				matrix[i+1][j-1] = 1
				matrix[i+1][j+1] = 1
			}
		}
	}

	return sum
}

func Part2(input string) int {
	matrix := Parse(input)
	sum := 0

	for i := 0; i < len(matrix)-1; i++ {
		line := matrix[i]
		for j, val := range line {
			if val <= 0 {
				continue
			}

			below := matrix[i+1][j]
			switch {
			case below >= 0:
				matrix[i+1][j] += val
			case below == SPLITTLER:
				matrix[i+1][j-1] += val
				matrix[i+1][j+1] += val
			}
		}
	}

	for _, val := range matrix[len(matrix)-1] {
		if val > 0 {
			sum += val
		}
	}

	return sum
}

func main() {
	fmt.Println("2025 day 07 solution")

	start := time.Now()
	fmt.Println("* Part1:")
	fmt.Println("\t* Result: ", Part1(input))
	fmt.Println("\t* Time: ", time.Since(start))

	start = time.Now()
	fmt.Println("* Part2:")
	fmt.Println("\t* Result: ", Part2(input))
	fmt.Println("\t* Time: ", time.Since(start))
}
