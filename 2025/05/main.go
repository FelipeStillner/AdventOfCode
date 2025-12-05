package main

import (
	_ "embed"
	"fmt"
	"sort"
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

func Parse(input string) ([]Range, []int) {
	var ranges []Range
	var nums []int
	input = strings.TrimSpace(input)
	before, after, _ := strings.Cut(input, "\n\n")
	before_lines := strings.SplitSeq(before, "\n")
	for line := range before_lines {
		b, a, _ := strings.Cut(line, "-")
		begin, _ := strconv.Atoi(b)
		end, _ := strconv.Atoi(a)
		ranges = append(ranges, Range{begin, end})
	}

	after_lines := strings.SplitSeq(after, "\n")
	for line := range after_lines {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}
	return ranges, nums
}

func Part1(input string) int {
	ranges, nums := Parse(input)
	sum := 0

	for _, num := range nums {
		for _, r := range ranges {
			if num >= r.begin && num <= r.end {
				sum++
				break
			}
		}
	}

	return sum
}

func Part2(input string) int {
	ranges, _ := Parse(input)
	sum := 0
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].begin < ranges[j].begin
	})
	val := 0

	for _, r := range ranges {
		init := max(val, r.begin)
		if r.end >= init {
			sum += r.end - init + 1
			val = r.end + 1
		}
	}

	return sum
}

func main() {
	fmt.Println("2025 day 05 solution")

	start := time.Now()
	fmt.Println("* Part1:")
	fmt.Println("\t* Result: ", Part1(input))
	fmt.Println("\t* Time: ", time.Since(start))

	start = time.Now()
	fmt.Println("* Part2:")
	fmt.Println("\t* Result: ", Part2(input))
	fmt.Println("\t* Time: ", time.Since(start))
}
