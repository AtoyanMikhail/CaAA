package main

import (
	"fmt"
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	testCases := []struct {
		size int
	}{
		{2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10},
		{11}, {12}, {13}, {14}, {15}, {16}, {17}, {18}, {19}, {20},
	}
	
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("Size:%d", tc.size), func(b *testing.B) {
			b.StopTimer() 
			t := New(tc.size)
			b.StartTimer() 

			t.Solve()
		})
	}
}
