package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func Parse(input string) [][]int {
	var nums [][]int
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var lineNums []int
		for _, c := range line {
			if c == '.' {
				lineNums = append(lineNums, 0)
			}
			if c == '@' {
				lineNums = append(lineNums, 1)
			}
		}
		nums = append(nums, lineNums)
	}
	return nums
}

func countNeighbors(lines [][]int, k, l int) int {
	count := 0
	if k > 0 && lines[k-1][l] == 1 {
		count++
	}
	if k < len(lines)-1 && lines[k+1][l] == 1 {
		count++
	}
	if l > 0 && lines[k][l-1] == 1 {
		count++
	}
	if l < len(lines[0])-1 && lines[k][l+1] == 1 {
		count++
	}
	if k > 0 && l > 0 && lines[k-1][l-1] == 1 {
		count++
	}
	if k > 0 && l < len(lines[0])-1 && lines[k-1][l+1] == 1 {
		count++
	}
	if k < len(lines)-1 && l > 0 && lines[k+1][l-1] == 1 {
		count++
	}
	if k < len(lines)-1 && l < len(lines[0])-1 && lines[k+1][l+1] == 1 {
		count++
	}
	return count
}

func removeRoll(lines [][]int) int {
	sum := 0
	for k, line := range lines {
		for l, c := range line {
			if c == 1 {
				neighbors := countNeighbors(lines, k, l)
				if neighbors < 4 {
					sum++
					lines[k][l] = 0
				}
			}
		}
	}
	return sum
}

func Part1(input string) int {
	lines := Parse(input)
	sum := 0

	for k, line := range lines {
		for l, c := range line {
			if c == 1 {
				neighbors := countNeighbors(lines, k, l)
				if neighbors < 4 {
					sum++
				}
			}
		}
	}

	return sum
}

func Part2(input string) int {
	lines := Parse(input)
	sum := 0
	removed := 1

	for removed > 0 {
		removed = removeRoll(lines)
		sum += removed
	}

	return sum
}

func main() {
	fmt.Println("2025 day 04 solution")

	start := time.Now()
	fmt.Println("* Part1:")
	fmt.Println("\t* Result: ", Part1(input))
	fmt.Println("\t* Time: ", time.Since(start))

	start = time.Now()
	fmt.Println("* Part2:")
	fmt.Println("\t* Result: ", Part2(input))
	fmt.Println("\t* Time: ", time.Since(start))
}
