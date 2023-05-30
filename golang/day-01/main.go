package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) int {
	total := 0
	current := 0
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			if current > total {
				total = current
			}
			current = 0
		} else {
			calori, _ := strconv.Atoi(input)
			current += calori
		}
	}
	return total
}

func part2(scanner *bufio.Scanner) int {
	calories := make([]int, 0, 10)
	current := 0
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			calories = append(calories, current)
			current = 0
		} else {
			calori, _ := strconv.Atoi(input)
			current += calori
		}
	}
	if current > 0 {
		calories = append(calories, current)
	}
	sort.Ints(calories)
	size := len(calories)

	return calories[size-1] + calories[size-2] + calories[size-3]
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// for scanner.Scan() {
	// 	input := strings.TrimSpace(scanner.Text())
	// 	if input == "" {
	// 		calories = append(calories, currentCalories)
	// 		currentCalories = 0
	// 		continue
	// 	}
	//
	// 	calori, _ := strconv.Atoi(input)
	// 	currentCalories += calori
	// }
	//
	// sort.Ints(calories)
	// size := len(calories)
	//
	// total := calories[size-1] + calories[size-2] + calories[size-3]
	//
	// log.Printf("max calori: %d\n", calories[size-1])
	// log.Printf("sum of top three calories: %d\n", total)
	//
	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}
