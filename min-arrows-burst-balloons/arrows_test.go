package arrows

import (
	"testing"
)

type TestCase struct {
	Input  [][]int
	Result int
}

func TestMinArrows(t *testing.T) {
	testCases := []TestCase{
		{
			Input:  [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			Result: 2,
		},
		{
			Input:  [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			Result: 4,
		},
		{
			Input:  [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			Result: 2,
		},
		{
			Input:  [][]int{{-2147483648, 2147483647}},
			Result: 1,
		},
	}
	for _, c := range testCases {
		result := MinArrows(c.Input)
		if c.Result != result {
			t.Fatalf("For %v: Expected %v, found %v", c.Input, c.Result, result)
		}
	}
}
