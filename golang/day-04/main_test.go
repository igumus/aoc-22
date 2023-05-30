package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay04(t *testing.T) {
	testcases := []struct {
		name  string
		path  string
		full  bool
		count int
	}{
		{
			name:  "test-part1",
			path:  "./test.txt",
			full:  true,
			count: 2,
		},
		{
			name:  "input-part1",
			path:  "./input.txt",
			full:  true,
			count: 483,
		},
		{
			name:  "test-part2",
			path:  "./test.txt",
			count: 4,
		},
		{
			name:  "input-part2",
			path:  "./input.txt",
			count: 874,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.path)
			require.Nil(t, err)
			defer file.Close()
			scanner := bufio.NewScanner(file)
			require.Equal(t, tc.count, countOverlap(scanner, tc.full))
		})
	}
}
