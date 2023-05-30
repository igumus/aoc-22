package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay05(t *testing.T) {
	testcases := []struct {
		name    string
		path    string
		f       func(*bufio.Scanner, bool) string
		reverse bool
		result  string
	}{
		{
			name:    "test-part1",
			path:    "./test.txt",
			f:       calculate,
			reverse: true,
			result:  "CMZ",
		},
		{
			name:    "input-part1",
			path:    "./input.txt",
			reverse: true,
			f:       calculate,
			result:  "HBTMTBSDC",
		},
		{
			name:   "test-part2",
			path:   "./test.txt",
			f:      calculate,
			result: "MCD",
		},
		{
			name:   "input-part2",
			path:   "./input.txt",
			f:      calculate,
			result: "PQTJRSHWS",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.path)
			require.Nil(t, err)
			defer file.Close()
			scanner := bufio.NewScanner(file)
			require.Equal(t, tc.result, tc.f(scanner, tc.reverse))
		})
	}
}
