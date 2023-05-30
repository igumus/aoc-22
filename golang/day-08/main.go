package main

import (
	"bufio"
	"strconv"
)

func parseGrid(sc *bufio.Scanner) [][]int {
	board := make([][]int, 0)
	for sc.Scan() {
		input := sc.Text()
		rowCount := len(input)
		row := make([]int, 0, rowCount)
		for i := 0; i < len(input); i++ {
			val, _ := strconv.Atoi(string(input[i]))
			row = append(row, val)
		}
		board = append(board, row)
	}
	return board
}

func calculateScenicScore(board [][]int) int {
	maxScore := 0
	for x := 1; x < len(board)-1; x++ {
		for y := 1; y < len(board[0])-1; y++ {
			curr := board[x][y]

			upScore := 0
			for xx := x - 1; xx >= 0; xx-- {
				upScore++
				if curr <= board[xx][y] {
					break
				}
			}

			downScore := 0
			for xx := x + 1; xx < len(board); xx++ {
				downScore++
				if curr <= board[xx][y] {
					break
				}
			}

			leftScore := 0
			for yy := y - 1; yy >= 0; yy-- {
				leftScore++
				if curr <= board[x][yy] {
					break
				}
			}

			rightScore := 0
			for yy := y + 1; yy < len(board[0]); yy++ {
				rightScore++
				if curr <= board[x][yy] {
					break
				}
			}
			score := upScore * leftScore * rightScore * downScore
			if maxScore < score {
				maxScore = score
			}

		}
	}
	return maxScore
}

func calculateVisible(board [][]int) int {
	count := (len(board)*2 + len(board[0])*2) - 4
	for x := 1; x < len(board)-1; x++ {
		for y := 1; y < len(board[0])-1; y++ {
			curr := board[x][y]

			upwardsVisible := curr > board[x-1][y]
			if upwardsVisible { // should i walk upwards
				for xx := x - 1; xx >= 0; xx-- {
					if curr <= board[xx][y] {
						upwardsVisible = false
						break
					}
				}
			}
			downwardsVisible := curr > board[x+1][y]
			if downwardsVisible { // should i walk downwards
				for xx := x + 1; xx < len(board); xx++ {
					if curr <= board[xx][y] {
						downwardsVisible = false
						break
					}
				}
			}

			leftwardsVisible := curr > board[x][y-1]
			if leftwardsVisible { // should i walk leftwards
				for yy := y - 1; yy >= 0; yy-- {
					if curr <= board[x][yy] {
						leftwardsVisible = false
						break
					}
				}
			}

			rightwardsVisible := curr > board[x][y+1]
			if rightwardsVisible { // should i walk rightwards
				for yy := y + 1; yy < len(board[0]); yy++ {
					if curr <= board[x][yy] {
						rightwardsVisible = false
						break
					}
				}
			}

			visible := upwardsVisible || downwardsVisible || rightwardsVisible || leftwardsVisible
			if visible {
				count += 1
			}
		}
	}
	return count
}
