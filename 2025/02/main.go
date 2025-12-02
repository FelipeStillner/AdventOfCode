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

type Range struct {
	begin int
	end   int
}

func Parse(input string) []Range {
	var ranges []Range
	input = strings.TrimSpace(input)
	lines := strings.Split(input, ",")
	for _, line := range lines {
		before, after, _ := strings.Cut(line, "-")
		begin, _ := strconv.Atoi(before)
		end, _ := strconv.Atoi(after)
		ranges = append(ranges, Range{begin, end})
	}
	return ranges
}

func getNextHalfNumber(n int) int {
	size := int(math.Floor(math.Log10(float64(n))) + 1)
	if size%2 == 0 {
		halfSize := int(size / 2)
		h1 := n / int(math.Pow10(halfSize))
		h2 := n % int(math.Pow10(halfSize))
		if h1 >= h2 {
			return h1
		}
		return h1 + 1
	}
	halfSize := size / 2
	return int(math.Pow10(halfSize))
}

func getLastHalfNumber(n int) int {
	size := int(math.Floor(math.Log10(float64(n))) + 1)
	if size%2 == 0 {
		halfSize := int(size / 2)
		h1 := n / int(math.Pow10(halfSize))
		h2 := n % int(math.Pow10(halfSize))
		if h1 <= h2 {
			return h1
		}
		return h1 - 1
	}
	halfSize := size / 2
	return int(math.Pow10(halfSize)) - 1
}

func doubleNumber(halfNum int) int {
	size := int(math.Floor(math.Log10(float64(halfNum))) + 1)
	doubleNum := halfNum * (1 + int(math.Pow10(size)))
	return doubleNum
}

func isRepetitionOf(n int64, r int64) bool {
	sizeR := int64(math.Floor(math.Log10(float64(r))) + 1)
	for n > 0 {
		if n%int64(math.Pow10(int(sizeR))) != r {
			return false
		}
		n /= int64(math.Pow10(int(sizeR)))
	}
	return true
}

func isRepetition(n int64) bool {
	size := int64(math.Floor(math.Log10(float64(n))) + 1)
	var halfSize int64 = size / 2
	for i := int64(1); i <= halfSize; i++ {
		power := int64(math.Pow10(int(i)))
		partialNum := n % power
		if partialNum == 0 {
			continue
		}
		if isRepetitionOf(n, partialNum) {
			return true
		}
	}
	return false
}

func Part1(input string) int {
	ranges := Parse(input)
	sum := 0

	for _, r := range ranges {
		halfBegin := getNextHalfNumber(r.begin)
		halfEnd := getLastHalfNumber(r.end)
		for i := halfBegin; i <= halfEnd; i++ {
			sum += doubleNumber(i)
		}
	}

	return sum
}

func Part2(input string) int64 {
	ranges := Parse(input)
	var sum int64 = 0
	for _, r := range ranges {
		for i := int64(r.begin); i <= int64(r.end); i++ {
			if isRepetition(i) {
				sum += i
			}
		}
	}
	return sum
}

func main() {
	fmt.Println("2025 day 02 solution")

	start := time.Now()
	fmt.Println("* Part1:")
	fmt.Println("\t* Result: ", Part1(input))
	fmt.Println("\t* Time: ", time.Since(start))

	start = time.Now()
	fmt.Println("* Part2:")
	fmt.Println("\t* Result: ", Part2(input))
	fmt.Println("\t* Time: ", time.Since(start))
}
