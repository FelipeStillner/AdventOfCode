package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Card struct {
	winning_numbers []int
	my_numbers      []int
	quantity        int
}

//go:embed input04.txt
var input string

func Parse(input string) []Card {
	var cards []Card
	trimmed_input := strings.TrimSpace(input)
	var lines = strings.Split(trimmed_input, "\n")
	for _, line := range lines {
		_, after, _ := strings.Cut(line, ":")
		winning, my, _ := strings.Cut(after, "|")
		w_str := strings.Fields(winning)
		m_str := strings.Fields(my)
		w_ints := make([]int, len(w_str))
		m_ints := make([]int, len(m_str))
		for i, val := range w_str {
			w_ints[i], _ = strconv.Atoi(val)
		}
		for i, val := range m_str {
			m_ints[i], _ = strconv.Atoi(val)
		}
		cards = append(cards, Card{
			winning_numbers: w_ints,
			my_numbers:      m_ints,
			quantity:        1,
		})
	}
	return cards
}

func Part1(input string) int {
	cards := Parse(input)
	sum := 0
	for _, card := range cards {
		points := 0
		for _, val := range card.my_numbers {
			if slices.Contains(card.winning_numbers, val) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		sum += points
	}
	return sum
}

func Part2(input string) int {
	cards := Parse(input)
	sum := 0
	for i, card := range cards {
		matching := 0
		for _, val := range card.my_numbers {
			if slices.Contains(card.winning_numbers, val) {
				matching++
			}
		}
		for j := range matching {
			if i+j+1 < len(cards) {
				cards[i+j+1].quantity += card.quantity
			}
		}
		sum += card.quantity
	}
	return sum
}

func main() {
	fmt.Println("2024 day 04 solution")

	start := time.Now()
	fmt.Println("* Part1:")
	fmt.Println("\t* Result: ", Part1(input))
	fmt.Println("\t* Time: ", time.Since(start))

	start = time.Now()
	fmt.Println("* Part2:")
	fmt.Println("\t* Result: ", Part2(input))
	fmt.Println("\t* Time: ", time.Since(start))
}
