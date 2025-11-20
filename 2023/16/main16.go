package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

const (
	TOP = iota
	BOTTOM
	LEFT
	RIGHT
)

const (
	NOTHING = iota
	SPLITTER_HORIZONTAL
	SPLITTER_VERTICAL
	MIRROR
	BACK_MIRROR
)

//go:embed input16.txt
var input string

func TranslateLayout(c rune) int {
	switch c {
	case '.':
		return NOTHING
	case '-':
		return SPLITTER_HORIZONTAL
	case '|':
		return SPLITTER_VERTICAL
	case '/':
		return MIRROR
	case '\\':
		return BACK_MIRROR
	default:
		panic("Error in Parsing")
	}
}

func Parse(input string) [][]int {
	trimmedInput := strings.TrimSuffix(input, "\n")
	var rows = strings.Split(trimmedInput, "\n")
	layout := make([][]int, len(rows))
	for i, row := range rows {
		for _, c := range row {
			layout[i] = append(layout[i], TranslateLayout(c))
		}
	}
	return layout
}

func SearchTile(searchMap [][][4]bool, layout [][]int, x int, y int, from int, lenx int, leny int) {
	if !(x >= 0 && y >= 0 && x < lenx && y < leny) {
		return
	}
	if searchMap[x][y][from] == true {
		return
	}
	searchMap[x][y][from] = true
	tile := layout[x][y]

	goLeft :=
		(tile == NOTHING && from == RIGHT) ||
			(tile == SPLITTER_HORIZONTAL && from != LEFT) ||
			(tile == MIRROR && from == TOP) ||
			(tile == BACK_MIRROR && from == BOTTOM)
	goRight :=
		(tile == NOTHING && from == LEFT) ||
			(tile == SPLITTER_HORIZONTAL && from != RIGHT) ||
			(tile == MIRROR && from == BOTTOM) ||
			(tile == BACK_MIRROR && from == TOP)
	goTop :=
		(tile == NOTHING && from == BOTTOM) ||
			(tile == SPLITTER_VERTICAL && from != TOP) ||
			(tile == MIRROR && from == LEFT) ||
			(tile == BACK_MIRROR && from == RIGHT)
	goBottom :=
		(tile == NOTHING && from == TOP) ||
			(tile == SPLITTER_VERTICAL && from != BOTTOM) ||
			(tile == MIRROR && from == RIGHT) ||
			(tile == BACK_MIRROR && from == LEFT)

	if goLeft {
		SearchTile(searchMap, layout, x, y-1, RIGHT, lenx, leny)
	}
	if goRight {
		SearchTile(searchMap, layout, x, y+1, LEFT, lenx, leny)
	}
	if goTop {
		SearchTile(searchMap, layout, x-1, y, BOTTOM, lenx, leny)
	}
	if goBottom {
		SearchTile(searchMap, layout, x+1, y, TOP, lenx, leny)
	}
}

func ExecuteSearch(layout [][]int, x int, y int, from int) int {
	lenx := len(layout)
	leny := len(layout[0])
	searchMap := make([][][4]bool, lenx)
	for i := range lenx {
		searchMap[i] = make([][4]bool, leny)
	}
	SearchTile(searchMap, layout, x, y, from, lenx, leny)
	sum := 0
	for _, v1 := range searchMap {
		for _, v2 := range v1 {
			if v2[0] || v2[1] || v2[2] || v2[3] {
				sum++
			}
		}
	}
	return sum
}

func Part1(input string) int {
	layout := Parse(input)
	return ExecuteSearch(layout, 0, 0, LEFT)
}

func Part2(input string) int {
	layout := Parse(input)
	lenx := len(layout)
	leny := len(layout[0])
	res := 0
	for i := range lenx {
		res = max(ExecuteSearch(layout, i, 0, LEFT), res)
		res = max(ExecuteSearch(layout, i, leny-1, RIGHT), res)
	}
	for i := range leny {
		res = max(ExecuteSearch(layout, 0, i, TOP), res)
		res = max(ExecuteSearch(layout, lenx-1, i, BOTTOM), res)
	}
	return res
}

func main() {
	fmt.Println("2023 day 16 solution")

	start := time.Now()
	fmt.Println("* Part1:")
	fmt.Println("\t* Result: ", Part1(input))
	fmt.Println("\t* Time: ", time.Since(start))

	start = time.Now()
	fmt.Println("* Part2:")
	fmt.Println("\t* Result: ", Part2(input))
	fmt.Println("\t* Time: ", time.Since(start))
}
