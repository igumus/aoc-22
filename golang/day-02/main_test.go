package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay02(t *testing.T) {
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
			count: 15,
		},
		{
			name:  "test-part2",
			path:  "./test.txt",
			f:     part2,
			count: 12,
		},
		{
			name:  "input-part1",
			path:  "./input.txt",
			f:     part1,
			count: 12679,
		},
		{
			name:  "input-part2",
			path:  "./input.txt",
			f:     part2,
			count: 14470,
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
