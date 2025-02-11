package main

type Square struct {
	X, Y, Size int
}

func (s Square) Overlap(other Square) bool {
	return s.X+s.Size > other.X && s.X < other.X+other.Size && s.Y+s.Size > other.Y && s.Y < other.Y+other.Size
}