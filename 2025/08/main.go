package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/FelipeStillner/AdventOfCode/2025/utils"
)

//go:embed input.txt
var input string

type Point struct {
	X, Y, Z int
}

type Line struct {
	Start, End int
	Distance   float64
}

func Parse(input string) []Point {
	var points []Point

	input = strings.TrimSpace(input)
	lines := strings.SplitSeq(input, "\n")

	for line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		points = append(points, Point{X: x, Y: y, Z: z})
	}

	return points
}

func getLines(points []Point) []Line {
	var lines []Line

	for i1, p1 := range points {
		for i2 := i1 + 1; i2 < len(points); i2++ {
			p2 := points[i2]
			dx := math.Pow(float64(p1.X-p2.X), 2)
			dy := math.Pow(float64(p1.Y-p2.Y), 2)
			dz := math.Pow(float64(p1.Z-p2.Z), 2)
			dist := math.Sqrt(dx + dy + dz)
			lines = append(lines, Line{Start: i1, End: i2, Distance: dist})
		}
	}

	return lines
}

func getNSmallerLines(lines []Line, n int) []Line {
	sort.Slice(lines, func(a, b int) bool {
		return lines[a].Distance < lines[b].Distance
	})

	var nearestLines []Line
	for i := range n {
		nearestLines = append(nearestLines, lines[i])
	}

	return nearestLines
}

func Part1(input string, nConnections int) int {
	points := Parse(input)

	lines := getLines(points)
	nSmallerLines := getNSmallerLines(lines, nConnections)

	unionFind := utils.NewUnionFind(len(points))
	for _, line := range nSmallerLines {
		unionFind.Union(line.Start, line.End)
	}

	sizes := unionFind.GetSetsNumberOfElements()

	res := sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]

	return res
}

func Part2(input string) int {
	points := Parse(input)
	res := 0

	lines := getLines(points)

	sort.Slice(lines, func(a, b int) bool {
		return lines[a].Distance < lines[b].Distance
	})

	unionFind := utils.NewUnionFind(len(points))
	for _, line := range lines {
		unionFind.Union(line.Start, line.End)
		if unionFind.Count == 1 {
			x1 := points[line.Start].X
			x2 := points[line.End].X
			res = x1 * x2
			break
		}
	}

	return res
}

func main() {
	fmt.Println("2025 day 08 solution")

	start := time.Now()
	fmt.Println("* Part1:")
	fmt.Println("\t* Result: ", Part1(input, 1000))
	fmt.Println("\t* Time: ", time.Since(start))

	start = time.Now()
	fmt.Println("* Part2:")
	fmt.Println("\t* Result: ", Part2(input))
	fmt.Println("\t* Time: ", time.Since(start))
}
