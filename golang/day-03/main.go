package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const (
	SIZE            = 53
	BOUND_LOWERCASE = 96
	BOUND_UPPERCASE = 38
)

// Clears the bit at pos in n.
// func clearBit(n int, pos uint) int {
// 	mask := ^(1 << pos)
// 	n &= mask
// 	return n
// }

// func hasBit(n int, pos uint) bool {
// 	val := n & (1 << pos)
// 	return (val > 0)
// }

// Sets the bit at pos in the integer n.
func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

func toPosition(n int) int {
	if n <= 0 {
		return 0
	}
	for pos := 0; pos < SIZE; pos++ {
		val := n & (1 << pos)
		if val > 0 {
			return pos
		}
	}

	return 0
}

func decode(items ...string) []int {
	rets := make([]int, len(items))
	for j := 0; j < len(items); j++ {
		s := items[j]
		var (
			ret int  = 0
			val uint = 0
		)
		for i := 0; i < len(s); i++ {
			if unicode.IsLower(rune(s[i])) {
				val = uint(s[i] - BOUND_LOWERCASE)
			} else {
				val = uint(s[i] - BOUND_UPPERCASE)
			}
			ret = setBit(ret, val)
		}
		rets[j] = ret
	}
	return rets
}

func match(items ...int) int {
	size := len(items)
	if size == 0 {
		return 0
	}

	ret := items[0]
	for i := 1; i < size; i++ {
		ret &= items[i]
	}
	return toPosition(ret)
}

func part1(scanner *bufio.Scanner) int {
	total := 0
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		size := len(input)
		parts := decode(input[:size/2], input[size/2:])
		ret := match(parts...)
		total += ret
	}
	return total
}

func part2(scanner *bufio.Scanner) int {
	total := 0
	for scanner.Scan() {

		first := scanner.Text()
		scanner.Scan()
		second := scanner.Text()
		scanner.Scan()
		third := scanner.Text()
		parts := decode(first, second, third)
		ret := match(parts...)
		total += ret
	}
	return total
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println(part2(scanner))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
