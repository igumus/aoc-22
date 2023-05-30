package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay08(t *testing.T) {
	testcases := []struct {
		name   string
		path   string
		f      func([][]int) int
		result int
	}{
		{
			name:   "test-2",
			path:   "./test.txt",
			f:      calculateScenicScore,
			result: 8,
		},
		{
			name:   "test-1",
			path:   "./test.txt",
			f:      calculateVisible,
			result: 21,
		},
		{
			name:   "input-1",
			path:   "./input.txt",
			f:      calculateVisible,
			result: 1859,
		},
		{
			name:   "input-2",
			path:   "./input.txt",
			f:      calculateScenicScore,
			result: 332640,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.path)
			require.Nil(t, err)
			defer file.Close()
			sc := bufio.NewScanner(file)
			board := parseGrid(sc)
			require.Equal(t, tc.result, tc.f(board))
			require.Nil(t, sc.Err())
		})
	}
}
