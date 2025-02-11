package main

import "fmt"

type Result struct {
	Count   int
	Squares []Square
}

func (r Result) String() string {
	res := fmt.Sprintf("%d\n", r.Count)

	for _, sq := range r.Squares {
		res += fmt.Sprintf("%d %d %d\n", sq.X+1, sq.Y+1, sq.Size)
	}

	return res
}
