package main

import (
	"bufio"
	"strconv"
	"strings"
)

type Direction byte

const (
	Right Direction = iota
	Left
	Down
	Up
)

type Move struct {
	direction Direction
	amount    int
}

type point struct{ x, y int }

func (p point) isConnected(o point) bool {
	if p.equals(o) {
		return true
	}
	if p.x == o.x {
		return (p.y+1) == o.y || (p.y-1) == o.y
	}

	if p.y == o.y {
		return (p.x+1) == o.x || (p.x-1) == o.x
	}

	if (p.x-1) == o.x && (p.y+1) == o.y {
		return true
	}
	if (p.x-1) == o.x && (p.y-1) == o.y {
		return true
	}
	if (p.x+1) == o.x && (p.y-1) == o.y {
		return true
	}
	if (p.x+1) == o.x && (p.y+1) == o.y {
		return true
	}
	return false
}

func (p point) equals(o point) bool {
	return p.x == o.x && p.y == o.y
}

func parseMoves(sc *bufio.Scanner) []Move {
	moves := make([]Move, 0)
	for sc.Scan() {
		input := strings.Fields(sc.Text())
		amount, _ := strconv.Atoi(input[1])
		move := Move{amount: amount}
		switch input[0] {
		case "R":
			move.direction = Right
			break
		case "L":
			move.direction = Left
			break
		case "U":
			move.direction = Up
			break
		case "D":
			move.direction = Down
			break
		}
		moves = append(moves, move)
	}
	return moves
}

func moveTail(direction Direction, head point, tail point) point {
	newTail := tail
	if !newTail.isConnected(head) {
		if head.y < tail.y {
			newTail.y -= 1
		} else if head.y > tail.y {
			newTail.y += 1
		}

		if head.x > tail.x {
			newTail.x += 1
		} else if head.x < tail.x {
			newTail.x -= 1
		}
		return newTail
	}
	return newTail
}

func calculate(moves []Move, tailsLen int) int {
	knots := make([]point, tailsLen)
	for i := 0; i < tailsLen; i++ {
		knots = append(knots, point{x: 0, y: 0})
	}
	pos := make(map[point]struct{})
	pos[knots[tailsLen-1]] = struct{}{}

	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			switch move.direction {
			case Right:
				knots[0].x += 1
				break
			case Left:
				knots[0].x -= 1
				break
			case Up:
				knots[0].y += 1
				break
			case Down:
				knots[0].y -= 1
				break
			}
			for j := range knots[:tailsLen-1] {
				knots[j+1] = moveTail(move.direction, knots[j], knots[j+1])
			}
			pos[knots[tailsLen-1]] = struct{}{}
		}
	}
	return len(pos)
}
