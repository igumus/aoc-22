package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMatch(t *testing.T) {
	vals := decode("QlJBfZssjgZsQs", "DNbScdDZdNcvvF")
	val1 := vals[0]
	val2 := vals[1]
	ret := val1 & val2
	assert.Equal(t, 52, toPosition(ret))
}

func TestDay03(t *testing.T) {
	testcases := []struct {
		name  string
		path  string
		f     func(*bufio.Scanner) int
		count int
	}{
		{
			name:  "test-part1",
			path:  "./test.txt",
			f:     part1,
			count: 157,
		},
		{
			name:  "input-part1",
			path:  "./input.txt",
			f:     part1,
			count: 8349,
		},
		{
			name:  "test-part2",
			path:  "./test.txt",
			f:     part2,
			count: 70,
		},
		{
			name:  "input-part2",
			path:  "./input.txt",
			f:     part2,
			count: 2681,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.path)
			require.Nil(t, err)
			defer file.Close()
			scanner := bufio.NewScanner(file)
			require.Equal(t, tc.count, tc.f(scanner))
		})
	}
}
