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

func Parse(input string) [][][3]int {
	trimmedInput := strings.TrimSuffix(input, "\n")
	games := strings.Split(trimmedInput, "\n")
	var list [][][3]int

	for _, game := range games {
		var gameList [][3]int
		_, after, _ := strings.Cut(game, ":")
		sets := strings.Split(after, ";")

		for _, set := range sets {
			var setList [3]int
			var colors = strings.Split(set, ",")

			for _, color := range colors {
				fields := strings.Fields(color)

				q, _ := strconv.Atoi(fields[0])
				c := fields[1]

				switch c {
				case "red":
					setList[0] = q
				case "green":
					setList[1] = q
				case "blue":
					setList[2] = q
				}
			}
			gameList = append(gameList, setList)
		}
		list = append(list, gameList)
	}
	return list
}

func Part1(input string) int {
	list := Parse(input)
	sum := 0

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	for i, game := range list {
		gameID := i + 1
		possible := true

		for _, set := range game {
			if set[0] > maxRed || set[1] > maxGreen || set[2] > maxBlue {
				possible = false
				break
			}
		}

		if possible {
			sum += gameID
		}
	}
	return sum
}

func Part2(input string) int {
	list := Parse(input)
	sum := 0

	for _, game := range list {
		maxRed := 0
		maxGreen := 0
		maxBlue := 0

		for _, set := range game {
			maxRed = max(set[0], maxRed)
			maxGreen = max(set[1], maxGreen)
			maxBlue = max(set[2], maxBlue)
		}

		sum += maxRed * maxGreen * maxBlue
	}
	return sum
}

func main() {
	fmt.Println("2023 day 02 solution")

	start := time.Now()
	fmt.Println("* Part1:")
	fmt.Println("\t* Result: ", Part1(input))
	fmt.Println("\t* Time: ", time.Since(start))

	start = time.Now()
	fmt.Println("* Part2:")
	fmt.Println("\t* Result: ", Part2(input))
	fmt.Println("\t* Time: ", time.Since(start))
}
