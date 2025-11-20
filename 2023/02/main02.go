package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input02.txt
var input string

func ParsePart1(input string) [][]map[string]int {
	var games = strings.Split(input, "\n")
	var list [][]map[string]int
	for _, game := range games {
		var gameList []map[string]int
		var _, after, _ = strings.Cut(game, ":")
		var sets = strings.Split(after, ";")
		for _, set := range sets {
			var setList map[string]int
			setList = make(map[string]int)
			var colors = strings.Split(set, ",")
			for _, color := range colors {
				var l, r, _ = strings.Cut(strings.Trim(color, " "), " ")
				q, _ := strconv.Atoi(strings.Trim(l, " "))
				c := strings.Trim(r, " ")
				setList[c] = q
			}
			gameList = append(gameList, setList)
		}
		list = append(list, gameList)
	}
	return list
}

func Part1(input string) int {
	list := ParsePart1(input)
	sum := 0
	for i, game := range list {
		for _, set := range game {
			if 12 < set["red"] || 13 < set["green"] || 14 < set["blue"] {
				sum += i + 1
				break
			}
		}
	}
	return sum
}

func Part2(input string) int {
	return 0
}

func main() {
	fmt.Println("2022 day 02 solution")

	start := time.Now()
	fmt.Println("* Part1:")
	fmt.Println("\t* Result: ", Part1(input))
	fmt.Println("\t* Time: ", time.Since(start))

	start = time.Now()
	fmt.Println("* Part2:")
	fmt.Println("\t* Result: ", Part2(input))
	fmt.Println("\t* Time: ", time.Since(start))
}
