package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay01(t *testing.T) {
	testcases := []struct {
		name  string
		path  string
		f     func(*bufio.Scanner) int
		count int
	}{
		{
			name:  "test-part1",
			path:  "./test.txt",
			count: 24000,
			f:     part1,
		},
		{
			name:  "input-part1",
			path:  "./input.txt",
			count: 71471,
			f:     part1,
		},
		{
			name:  "test-part2",
			path:  "./test.txt",
			count: 45000,
			f:     part2,
		},
		{
			name:  "input-part2",
			path:  "./input.txt",
			count: 211189,
			f:     part2,
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
