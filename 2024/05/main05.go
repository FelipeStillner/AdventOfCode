package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Conversion struct {
	dest_range_start   int
	source_range_start int
	len                int
}

//go:embed input05.txt
var input string

func Parse(input string) ([]int, [][]Conversion) {
	input = strings.TrimSpace(input)
	var seeds []int
	var conversions [][]Conversion

	seeds_a, convs_a, _ := strings.Cut(input, "\n")
	_, seeds_b, _ := strings.Cut(seeds_a, ":")
	seeds_c := strings.FieldsSeq(seeds_b)
	for seed_str := range seeds_c {
		seeds_int, _ := strconv.Atoi(seed_str)
		seeds = append(seeds, seeds_int)
	}

	convs_b := strings.SplitSeq(convs_a, "\n\n")
	for conv_b := range convs_b {
		deep_conversions := make([]Conversion, 0)
		conv_b = strings.TrimSpace(conv_b)
		_, conv_c, _ := strings.Cut(conv_b, ":\n")

		convs_d := strings.SplitSeq(conv_c, "\n")
		for conv_d := range convs_d {
			conv_str := strings.Fields(conv_d)
			dest, _ := strconv.Atoi(conv_str[0])
			source, _ := strconv.Atoi(conv_str[1])
			length, _ := strconv.Atoi(conv_str[2])
			deep_conversions = append(deep_conversions, Conversion{
				dest_range_start:   dest,
				source_range_start: source,
				len:                length,
			})
		}
		conversions = append(conversions, deep_conversions)
	}
	return seeds, conversions
}

func Part1(input string) int {
	values, conversions := Parse(input)
	for _, step := range conversions {
		for j, v := range values {
			for _, conv := range step {
				if v >= conv.source_range_start && v < conv.source_range_start+conv.len {
					values[j] = conv.dest_range_start + (v - conv.source_range_start)
				}
			}
		}
	}
	min := math.MaxInt32
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

type Range struct {
	start int
	size  int
}

func Part2(input string) int {
	values, conversions := Parse(input)
	ranges := make([]Range, 0)
	for i := 0; i < len(values); i += 2 {
		ranges = append(ranges, Range{start: values[i], size: values[i+1]})
	}

	for _, step := range conversions {
		new_ranges := make([]Range, 0)
		for _, r := range ranges {
			for _, conv := range step {
				begin := max(r.start, conv.source_range_start)
				end := min(r.start+r.size, conv.source_range_start+conv.len)
				if begin < end {
					size := end - begin
					new_ranges = append(new_ranges, Range{start: conv.dest_range_start + (begin - conv.source_range_start), size: size})
				}
			}
		}
		ranges = new_ranges
	}

	min := math.MaxInt32
	for _, r := range ranges {
		if r.start < min {
			min = r.start
		}
	}
	return min
}

func main() {
	fmt.Println("2024 day 05 solution")

	start := time.Now()
	fmt.Println("* Part1:")
	fmt.Println("\t* Result: ", Part1(input))
	fmt.Println("\t* Time: ", time.Since(start))

	start = time.Now()
	fmt.Println("* Part2:")
	fmt.Println("\t* Result: ", Part2(input))
	fmt.Println("\t* Time: ", time.Since(start))
}
