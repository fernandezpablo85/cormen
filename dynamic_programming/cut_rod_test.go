package dynamic_programming

import (
	"testing"
)

var prices = []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}

var testCases = []struct {
	size int
	best int
}{{4, 10}, {5, 13}, {6, 17}, {7, 18}, {8, 22}, {9, 25}, {10, 30}}

func TestBruteMaxValue(t *testing.T) {
	for _, test := range testCases {
		value := BruteMaxValue(test.size, prices)
		if value != test.best {
			t.Fatalf("got %d, expected %d", value, test.best)
		}
	}
}

func TestMemoizedMaxValue(t *testing.T) {
	for _, test := range testCases {
		value := MemoizedMaxValue(test.size, prices)
		if value != test.best {
			t.Fatalf("got %d, expected %d", value, test.best)
		}
	}
}

func TestBottomUpMaxValue(t *testing.T) {
	for _, test := range testCases {
		value := BottomUpMaxValue(test.size, prices)
		if value != test.best {
			t.Fatalf("got %d, expected %d", value, test.best)
		}
	}
}
