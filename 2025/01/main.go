package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var input string

func Parse(input string) {
}

func Part1(input string) int {
	return 0
}

func Part2(input string) int {
	return 0
}

func main() {
	fmt.Println("2025 day 01 solution")

	start := time.Now()
	fmt.Println("* Part1:")
	fmt.Println("\t* Result: ", Part1(input))
	fmt.Println("\t* Time: ", time.Since(start))

	start = time.Now()
	fmt.Println("* Part2:")
	fmt.Println("\t* Result: ", Part2(input))
	fmt.Println("\t* Time: ", time.Since(start))
}
