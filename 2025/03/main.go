package main

import (
	_ "embed"
	"fmt"
	"strconv"
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
		for i := range line {
			num, _ := strconv.Atoi(string(line[i]))
			lineNums = append(lineNums, num)
		}
		nums = append(nums, lineNums)
	}
	return nums
}

func getBiggestValue(line []int, i_init int, i_end int) (int, int) {
	val := 0
	idx := 0
	for i := i_init; i < i_end; i++ {
		if line[i] > val {
			val = line[i]
			idx = i
		}
	}
	return val, idx
}

func Part1(input string) int {
	lines := Parse(input)
	sum := 0

	for _, line := range lines {
		tens, i_tens := getBiggestValue(line, 0, len(line)-1)
		units, _ := getBiggestValue(line, i_tens+1, len(line))
		sum += tens*10 + units
	}

	return sum
}

func Part2(input string) int {
	lines := Parse(input)
	sum := 0

	for _, line := range lines {
		value := 0
		i_init := 0

		for i := range 12 {
			val, i_val := getBiggestValue(line, i_init, len(line)-11+i)
			i_init = i_val + 1
			value = value*10 + val
		}

		sum += value
	}

	return sum
}

func main() {
	fmt.Println("2025 day 03 solution")

	start := time.Now()
	fmt.Println("* Part1:")
	fmt.Println("\t* Result: ", Part1(input))
	fmt.Println("\t* Time: ", time.Since(start))

	start = time.Now()
	fmt.Println("* Part2:")
	fmt.Println("\t* Result: ", Part2(input))
	fmt.Println("\t* Time: ", time.Since(start))
}
