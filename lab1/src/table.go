package main

import (
	"fmt"
	"math"
)

type Table struct {
	edgeLength    int
	current       []Square
	result        []Square
	bestCount     int
	currentCount  int
	maxSquareEdge int
}

func NewTable(edge int) *Table {
	t := &Table{
		edgeLength:    edge,
		bestCount:     math.MaxInt32,
		maxSquareEdge: edge - 1,
	}
	return t
}

func (t *Table) placeSquare(square Square) error {
	if square.X+square.Size > t.edgeLength || square.Y+square.Size > t.edgeLength {
		return fmt.Errorf("Out of bounds placement")
	}

	for _, sq := range t.current {
		if square.Overlap(sq) {
			return fmt.Errorf("Square overlaps with existing square")
		}
	}

	t.current = append(t.current, square)
	t.currentCount++
	return nil
}

func (t *Table) findEmptyX(y int) int {
	for x := 0; x < t.edgeLength; x++ {
		overlapped := false
		for _, sq := range t.current {
			if sq.Overlap(Square{x, y, 1}) {
				overlapped = true
				break
			}
		}

		if !overlapped {
			return x
		}
	}
	return -1
}

func (t *Table) popSquare() error {
	if len(t.current) == 0 {
		return fmt.Errorf("No squares to pop")
	}

	t.current = t.current[:len(t.current)-1]
	t.currentCount--
	return nil
}

func (t *Table) backtrack(y int) {
	if t.currentCount >= t.bestCount {
		return
	}

	if y >= t.edgeLength {
		if t.currentCount < t.bestCount {
			t.bestCount = t.currentCount
			t.result = make([]Square, t.bestCount)
			copy(t.result, t.current)
		}
		return
	}

	x := t.findEmptyX(y)

	if x == -1 {
		t.backtrack(y + 1)
		return
	}

	maxSize := min(t.maxSquareEdge, t.edgeLength-x, t.edgeLength-y)
	for size := maxSize; size >= 1; size-- {
		sq := Square{x, y, size}
		if err := t.placeSquare(sq); err == nil {
			t.backtrack(y)

			t.popSquare()
		}
	}
}

func (t *Table) optimize() error {
	if t.edgeLength%2 == 0 {
		t.result = []Square{
			{0, 0, t.edgeLength / 2},
			{t.edgeLength / 2, 0, t.edgeLength / 2},
			{0, t.edgeLength / 2, t.edgeLength / 2},
			{t.edgeLength / 2, t.edgeLength / 2, t.edgeLength / 2},
		}
		t.bestCount = 4
		return nil
	}

	if isPowerOfTwoMinusOne(t.edgeLength) {
		base := (t.edgeLength + 1) / 2

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
				Square{t.edgeLength - indentation, t.edgeLength - squareSize - indentation, squareSize},
				Square{t.edgeLength - squareSize - indentation, t.edgeLength - indentation, squareSize},
				Square{t.edgeLength - indentation, t.edgeLength - indentation, squareSize},
			)
			t.bestCount += 3
			squareSize /= 2
			indentation += squareSize
		}

		return nil
	}

	return fmt.Errorf("Could not optimize calculations")
}

func (t *Table) PlaceSquares() Result {
	if err := t.optimize(); err == nil {
		result := Result{
			Count:   t.bestCount,
			Squares: t.result,
		}

		return result
	}

	if isPrime(t.edgeLength) {
		base := (t.edgeLength + 1) / 2
		t.placeSquare(Square{0, 0, base})
		t.placeSquare(Square{0, base, base - 1})
		t.placeSquare(Square{base, 0, base - 1})
		t.current = []Square{
			{0, 0, base},
			{0, base, base - 1},
			{base, 0, base - 1},
		}
		t.currentCount = 3
		t.maxSquareEdge = base - 1
	}

	t.backtrack(0)

	result := Result{
		Count:   t.bestCount,
		Squares: t.result,
	}

	return result
}
