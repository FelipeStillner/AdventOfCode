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

const (
	PROD = iota
	SUM
)

type Count struct {
	nums      []int64
	operation int
}

func ParsePart1(input string) []Count {
	var counts []Count

	input = strings.TrimSpace(input)

	var nums [][]int64
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines)-1; i++ {
		var line_nums []int64
		line := lines[i]
		line_fields := strings.Fields(line)
		for _, field := range line_fields {
			num, _ := strconv.Atoi(field)
			line_nums = append(line_nums, int64(num))
		}
		nums = append(nums, line_nums)
	}

	operation_fields := strings.Fields(lines[len(lines)-1])
	for i := range operation_fields {
		var ns []int64
		for j := 0; j < len(nums); j++ {
			ns = append(ns, nums[j][i])
		}
		if operation_fields[i] == "+" {
			counts = append(counts, Count{operation: SUM, nums: ns})
		} else {
			counts = append(counts, Count{operation: PROD, nums: ns})
		}
	}

	return counts
}

func calculateCounts(count Count) int64 {
	if count.operation == SUM {
		res := int64(0)
		for _, num := range count.nums {
			res += num
		}
		return res
	} else {
		res := int64(1)
		for _, num := range count.nums {
			res *= num
		}
		return res
	}
}

func Part1(input string) int64 {
	counts := ParsePart1(input)
	sum := int64(0)
	for _, count := range counts {
		sum += calculateCounts(count)
	}
	return sum
}

func updateLineSizes(lines_nums []string, operation_line string) ([]string, string) {
	max_len := len(operation_line)
	for _, v := range lines_nums {
		max_len = max(max_len, len(v))
	}
	for _ = range max_len - len(operation_line) {
		operation_line += " "
	}
	for i, v := range lines_nums {
		for _ = range max_len - len(v) {
			lines_nums[i] += " "
		}
	}
	return lines_nums, operation_line
}

func getColumnNum(lines_nums []string, i int) (int64, bool) {
	num := int64(0)
	foundDigit := false
	for j := range lines_nums {
		if i >= len(lines_nums[j]) {
			continue
		}

		n, err := strconv.Atoi(string(lines_nums[j][i]))
		if err == nil {
			num *= 10
			num += int64(n)
			foundDigit = true
		}
	}
	return num, foundDigit
}

func ParsePart2(input string) []Count {
	var counts []Count
	input = strings.TrimSpace(input)

	lines := strings.Split(input, "\n")
	lines_nums := lines[:len(lines)-1]
	operation_line := lines[len(lines)-1]

	lines_nums, operation_line = updateLineSizes(lines_nums, operation_line)

	actual_nums := []int64{}
	actual_operation := -1

	for i := range operation_line {
		switch operation_line[i] {
		case '+':
			actual_operation = SUM
		case '*':
			actual_operation = PROD
		}

		num, found := getColumnNum(lines_nums, i)

		if found {
			actual_nums = append(actual_nums, num)
		}

		if i+1 < len(operation_line) && operation_line[i+1] != ' ' {
			if len(actual_nums) > 0 {
				count := Count{operation: actual_operation, nums: actual_nums}
				counts = append(counts, count)
				actual_nums = []int64{}
				actual_operation = -1
			}
		}
	}

	if len(actual_nums) > 0 {
		count := Count{operation: actual_operation, nums: actual_nums}
		counts = append(counts, count)
	}

	return counts
}

func Part2(input string) int64 {
	counts := ParsePart2(input)
	sum := int64(0)
	for _, count := range counts {
		sum += calculateCounts(count)
	}
	return sum
}

func main() {
	fmt.Println("2025 day 06 solution")

	start := time.Now()
	fmt.Println("* Part1:")
	fmt.Println("\t* Result: ", Part1(input))
	fmt.Println("\t* Time: ", time.Since(start))

	start = time.Now()
	fmt.Println("* Part2:")
	fmt.Println("\t* Result: ", Part2(input))
	fmt.Println("\t* Time: ", time.Since(start))
}
