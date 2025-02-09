package main

import (
	"fmt"
	"math"
)

type Square struct {
	X, Y, Size int
}

type Table struct {
	edge          int
	matrix        [][]int
	current       []Square
	result        []Square
	bestCount     int
	currentCount  int
	maxSquareEdge int
}

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

func New(edge int) *Table {
	t := &Table{
		edge:          edge,
		bestCount:     math.MaxInt32,
		maxSquareEdge: edge - 1,
	}
	t.matrix = make([][]int, edge)
	for i := range t.matrix {
		t.matrix[i] = make([]int, edge)
	}
	return t
}

func (t *Table) Place(x, y, size int) error {
	if x+size > t.edge || y+size > t.edge {
		return fmt.Errorf("Out of bounds placement")
	}

	for i := y; i < y+size; i++ {
		for j := x; j < x+size; j++ {
			if t.matrix[i][j] != 0 {
				return fmt.Errorf("Could not place square at coordinates: x:%v y:%v", x, y)
			}
		}
	}

	for i := y; i < y+size; i++ {
		for j := x; j < x+size; j++ {
			t.matrix[i][j] = size
		}
	}
	return nil
}

func (t *Table) FindEmptyX(y int) int {
	for x := 0; x < t.edge; x++ {
		if t.matrix[y][x] == 0 {
			return x
		}
	}
	return -1
}

func (t *Table) RemoveSquare(x, y, size int) {
	for i := y; i < y+size; i++ {
		for j := x; j < x+size; j++ {
			t.matrix[i][j] = 0
		}
	}
}

func (t *Table) Backtrack(y int) {
	if y >= t.edge {
		if t.currentCount < t.bestCount {
			t.bestCount = t.currentCount
			t.result = make([]Square, t.bestCount)
			copy(t.result, t.current)
		}
		return
	}

	x := t.FindEmptyX(y)
	if x == -1 {
		t.Backtrack(y + 1)
		return
	}

	if t.currentCount >= t.bestCount {
		return
	}

	maxSize := min(t.maxSquareEdge, t.edge-x, t.edge-y)
	for size := maxSize; size >= 1; size-- {
		if err := t.Place(x, y, size); err == nil {
			t.current = append(t.current, Square{x, y, size})
			t.currentCount++

			t.Backtrack(y)

			t.RemoveSquare(x, y, size)
			t.current = t.current[:len(t.current)-1]
			t.currentCount--
		}
	}
}

func (t *Table) Optimize() error {
	if t.edge%2 == 0 {
		t.result = []Square{
			{0, 0, t.edge / 2},
			{t.edge / 2, 0, t.edge / 2},
			{0, t.edge / 2, t.edge / 2},
			{t.edge / 2, t.edge / 2, t.edge / 2},
		}
		t.bestCount = 4
		return nil
	}

	if isPowerOfTwoMinusOne(t.edge) {
		base := (t.edge + 1) / 2

		t.result = []Square{
			{0, 0, base},
			{0, base, base - 1},
			{base, 0, base - 1},
		}
		t.bestCount = 3

		squareSize := base / 2
		indentation := squareSize
		for squareSize > 0 {
			t.result = append(t.result,
				Square{t.edge - indentation, t.edge - squareSize - indentation, squareSize},
				Square{t.edge - squareSize - indentation, t.edge - indentation, squareSize},
				Square{t.edge - indentation, t.edge - indentation, squareSize},
			)
			t.bestCount += 3
			squareSize /= 2
			indentation += squareSize
		}

		return nil
	}

	return fmt.Errorf("Could not optimize calculations")
}

func (t *Table) FindSquares() Result {
	if err := t.Optimize(); err == nil {
		result := Result{
			Count:   t.bestCount,
			Squares: t.result,
		}

		return result
	}

	if isPrime(t.edge) {
		base := (t.edge + 1) / 2
		t.Place(0, 0, base)
		t.Place(0, base, base-1)
		t.Place(base, 0, base-1)
		t.current = []Square{
			{0, 0, base},
			{0, base, base - 1},
			{base, 0, base - 1},
		}
		t.currentCount = 3
		t.maxSquareEdge = base - 1
	}

	t.Backtrack(0)

	result := Result{
		Count:   t.bestCount,
		Squares: t.result,
	}

	return result
}
