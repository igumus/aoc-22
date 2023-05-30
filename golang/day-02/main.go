package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type item int

const (
	Unknown item = iota
	Rock
	Paper
	Scissors
)

const (
	Lost int = 0
	Draw int = 3
	Won  int = 6
)

// decodes input to item
func decode(move string) item {
	switch move {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}
	return Unknown
}

// Score calculation for single round;
// The shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) +
// the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won).
func scorePartOne(m item, o item) int {
	if m == o {
		return Draw + int(m)
	}

	if m == Rock {
		if o == Scissors {
			return Won + int(m) // i win
		}
		return Lost + int(m) // opponent wins
	}

	if m == Paper {
		if o == Rock {
			return Won + int(m) // i win
		}
		return Lost + int(m) // opponent wins
	}
	if m == Scissors {
		if o == Paper {
			return Won + int(m) // i win
		}
		return Lost + int(m) // opponent wins
	}
	return 0
}

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win.
func scorePartTwo(m item, o item) int {
	// need to draw
	if m == Paper {
		return Draw + int(o)
	}

	// need to lose
	if m == Rock {
		switch o {
		case Rock:
			return Lost + int(Scissors)
		case Paper:
			return Lost + int(Rock)
		case Scissors:
			return Lost + int(Paper)
		}
	}

	// need to win
	if m == Scissors {
		switch o {
		case Rock:
			return Won + int(Paper)
		case Paper:
			return Won + int(Scissors)
		case Scissors:
			return Won + int(Rock)
		}
	}

	// invalid score
	return 0
}

func part1(scanner *bufio.Scanner) int {
	score := 0
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		items := strings.Split(input, " ")
		opponentMove := decode(items[0])
		myMove := decode(items[1])
		score += scorePartOne(myMove, opponentMove)
	}
	return score
}

func part2(scanner *bufio.Scanner) int {
	score := 0
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		items := strings.Split(input, " ")
		opponentMove := decode(items[0])
		myMove := decode(items[1])
		score += scorePartTwo(myMove, opponentMove)
	}
	return score
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Printf("part1: %d\n", part1(scanner))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
