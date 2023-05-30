package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay07(t *testing.T) {
	testcases := []struct {
		name   string
		path   string
		limit  int64
		f      func(*bufio.Scanner) int64
		result int64
	}{

		{
			name:   "test-1",
			path:   "./test.txt",
			f:      calculate,
			result: 95437,
		},
		{
			name:   "test-2",
			path:   "./test.txt",
			f:      calculatePartTwo,
			result: 24933642,
		},
		{
			name:   "input-2",
			path:   "./input.txt",
			f:      calculatePartTwo,
			result: 2481982,
		},
		{
			name:   "input-1",
			path:   "./input.txt",
			f:      calculate,
			result: 1517599,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.path)
			require.Nil(t, err)
			defer file.Close()
			sc := bufio.NewScanner(file)
			require.Equal(t, tc.result, tc.f(sc))
			require.Nil(t, sc.Err())
		})
	}
}
