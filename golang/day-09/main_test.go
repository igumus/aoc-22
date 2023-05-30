package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDay09(t *testing.T) {
	testcases := []struct {
		name      string
		path      string
		f         func([]Move, int) int
		tailCount int
		result    int
	}{
		{
			name:      "test-1",
			path:      "./test.txt",
			f:         calculate,
			tailCount: 2,
			result:    13,
		},
		{
			name:      "test-2",
			path:      "./test-2.txt",
			f:         calculate,
			tailCount: 10,
			result:    36,
		},
		{
			name:      "input-1",
			path:      "./input.txt",
			f:         calculate,
			tailCount: 2,
			result:    6522,
		},
		{
			name:      "input-2",
			path:      "./input.txt",
			f:         calculate,
			tailCount: 10,
			result:    2717,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.path)
			require.Nil(t, err)
			defer file.Close()
			sc := bufio.NewScanner(file)
			moves := parseMoves(sc)
			require.Equal(t, tc.result, tc.f(moves, tc.tailCount))
			require.Nil(t, sc.Err())
		})
	}
}
