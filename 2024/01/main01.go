package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed input01.txt
var input string

func ParsePart1(input string) ([]int, []int) {
	var left, right []int
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		var before, after, _ = strings.Cut(line, " ")
		var l, _ = strconv.Atoi(before)
		var r, _ = strconv.Atoi(strings.Trim(after, " "))
		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}

func Part1(input string) int {
	left, right := ParsePart1(input)
	slices.Sort(right)
	slices.Sort(left)
	sum := 0
	for i := range right {
		partialSum := math.Abs(float64(right[i] - left[i]))
		sum += int(partialSum)
	}
	return sum
}

func ParsePart2(input string) ([]int, map[int]int) {
	var left []int
	right := make(map[int]int)
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		var before, after, _ = strings.Cut(line, " ")
		var l, _ = strconv.Atoi(before)
		var r, _ = strconv.Atoi(strings.Trim(after, " "))
		left = append(left, l)
		right[r] += 1
	}
	return left, right
}

func Part2(input string) int {
	left, right := ParsePart2(input)
	sum := 0
	for _, val := range left {
		partialSum := val * right[val]
		sum += partialSum
	}
	return sum
}

func main() {
	fmt.Println("2024 day 01 solution")

	start := time.Now()
	fmt.Println("* Part1:")
	fmt.Println("\t* Result: ", Part1(input))
	fmt.Println("\t* Time: ", time.Since(start))

	start = time.Now()
	fmt.Println("* Part2:")
	fmt.Println("\t* Result: ", Part2(input))
	fmt.Println("\t* Time: ", time.Since(start))
}
