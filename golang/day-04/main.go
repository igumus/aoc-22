package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func countOverlap(scanner *bufio.Scanner, full bool) int {
	count := 0
	for scanner.Scan() {
		input := scanner.Text()
		// fmt.Println(input)
		elves := strings.Split(input, ",")
		// fmt.Println(elves)
		firstElve := strings.Split(elves[0], "-")
		fes, _ := strconv.Atoi(firstElve[0])
		fee, _ := strconv.Atoi(firstElve[1])
		//fmt.Printf("first pair range: [%d, %d]\n", fes, fee)
		secondElve := strings.Split(elves[1], "-")
		ses, _ := strconv.Atoi(secondElve[0])
		see, _ := strconv.Atoi(secondElve[1])
		maxStart := int(math.Max(float64(fes), float64(ses)))
		minEnd := int(math.Min(float64(fee), float64(see)))
		if maxStart <= minEnd {
			if full {
				if (maxStart == fes) && (minEnd == fee) || (maxStart == ses && minEnd == see) {
					count++
				}
			} else {
				count++
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println(countOverlap(scanner, true))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
