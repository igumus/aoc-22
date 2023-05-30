package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type board struct {
	cols [][]string
}

func newBoard(b [][]string) *board {
	size := len(b[0])
	ret := &board{
		cols: make([][]string, size),
	}

	for i := len(b) - 1; i >= 0; i-- {
		row := b[i]
		for j := 0; j < size; j++ {
			if row[j] != "-" {
				ret.cols[j] = append(ret.cols[j], row[j])
			}
		}

	}
	return ret
}

func (b *board) result() string {
	ret := ""
	for i := 0; i < len(b.cols); i++ {
		size := len(b.cols[i])
		ret += b.cols[i][size-1]
	}
	return ret
}

func (b *board) applyAllAction(actions []string, reverseOrder bool) {
	for _, action := range actions {
		b.applyAction(action, reverseOrder)
	}
}

func (b *board) applyAction(input string, reverseOrder bool) {
	input = strings.ReplaceAll(input, "move ", "")
	input = strings.ReplaceAll(input, " from ", ",")
	input = strings.ReplaceAll(input, " to ", ",")
	input = strings.TrimSpace(input)
	action := strings.Split(input, ",")
	count, _ := strconv.Atoi(action[0])
	from, _ := strconv.Atoi(action[1])
	from -= 1
	to, _ := strconv.Atoi(action[2])
	to -= 1
	tempSize := len(b.cols[from])
	tempVals := b.cols[from][tempSize-count : tempSize]
	if tempSize-count == 0 {
		b.cols[from] = make([]string, 0)
	} else {
		b.cols[from] = b.cols[from][:tempSize-count]
	}
	if reverseOrder {
		for i := len(tempVals) - 1; i >= 0; i-- {
			b.cols[to] = append(b.cols[to], tempVals[i])
		}
		return
	}
	for i := 0; i < len(tempVals); i++ {
		b.cols[to] = append(b.cols[to], tempVals[i])
	}
}

func calculate(scanner *bufio.Scanner, reverseOrder bool) string {
	actionStarted := false
	boardParsed := make([][]string, 0)
	actions := make([]string, 0)
	for scanner.Scan() {
		input := scanner.Text()
		if strings.TrimSpace(input) == "" {
			actionStarted = true
			continue
		}
		if strings.HasPrefix(input, " 1") {
			// ignore position
			continue
		}
		if !actionStarted {
			if strings.HasPrefix(input, "   ") {
				input = strings.Replace(input, "   ", "[-]", 1)
			}
			input = strings.ReplaceAll(input, "]    ", "] [-]")
			input = strings.ReplaceAll(input, "-]    ", "-] [-]")
			input = strings.ReplaceAll(input, "] [", ",")
			input = input[1 : len(input)-1]
			elements := strings.Split(input, ",")
			boardParsed = append(boardParsed, elements)
		} else {
			actions = append(actions, input)
		}
	}
	board := newBoard(boardParsed)
	board.applyAllAction(actions, reverseOrder)
	return board.result()
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println(calculate(scanner, true))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
