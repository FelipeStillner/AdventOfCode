package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

func Parse(input string) []int {
	var numbers []int
	input = strings.TrimSpace(input)
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		var n, _ = strconv.Atoi(line[1:])
		if line[0] == 'L' {
			numbers = append(numbers, -n)
		} else {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func Part1(input string) int {
	numbers := Parse(input)
	now := 50
	n_zeros := 0

	for _, n := range numbers {
		now += n

		now = now % 100

		if now == 0 {
			n_zeros++
		}
	}

	return n_zeros
}

func Part2(input string) int {
	numbers := Parse(input)
	dial := 50
	n_zeros := 0

	for _, n := range numbers {
		n_zeros += int(math.Abs(float64(n / 100)))

		dial_ant := dial
		dial += n % 100

		if (dial_ant > 0 && dial < 0) || (dial_ant < 0 && dial > 0) || dial == 0 {
			n_zeros++
		}

		if dial >= 100 {
			dial -= 100
			n_zeros++
		}

		if dial <= -100 {
			dial += 100
			n_zeros++
		}
	}

	return n_zeros
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
