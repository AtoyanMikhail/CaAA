package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n < 2 || n > 40 {
		panic("Invalid size")
	}

	t := New(n)
	r := t.Solve()

	fmt.Print(r)
}
