package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func findDistinct(input string, capacity int) int {
	i := 0
start:
	for i < len(input)-capacity {
		subInput := input[i : i+capacity]
		lookup := make(map[int]struct{}, 0)
		for j := 0; j < len(subInput); j++ {
			if _, ok := lookup[int(subInput[j])]; ok {
				i++
				goto start
			}
			lookup[int(subInput[j])] = struct{}{}
		}
		return i + capacity
	}
	return 0
}
